package ami

//go:generate stringer -type=CommandType -output=strings.go

import (
	"bufio"
	"github.com/pkg/errors"
	"net"
	"regexp"

	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

type CommandType int

const (
	Action CommandType = iota + 1
	Response
	Event
)

type Pack interface {
	GetType() CommandType
}

type Command struct {
	Type    CommandType
	Headers map[string]interface{}
}

type CommandResponse struct {
	Response *Command
	Err      error
}

type AMI interface {
	ServerVersion() string
	WriteCommand(interface{}) <-chan CommandResponse
	WriteCommandSync(interface{}) (*Command, error)
}

type Config struct {
	Username  string
	Secret    string
	Listeners []chan<- *Command
}

func Dial(addr string, config Config) (AMI, error) {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	success := false
	var ami *ami
	defer func() {
		if !success {
			con.Close()
			if ami != nil {
				ami.Close(nil)
			}
		}
	}()

	scanner := bufio.NewScanner(con)
	ami = newAMI()
	if scanner.Scan() {
		version := scanner.Text()
		match := regexp.MustCompile("Asterisk Call Manager/([0-9.]+)").FindStringSubmatch(version)
		if len(match) > 1 {
			ami.serverVersion = match[1]
		} else {
			return nil, errors.Errorf("Invalid server version: %s", version)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	ami.scanner = scanner
	ami.con = con
	ami.subs = config.Listeners
	go ami.start()

	res, err := ami.WriteCommandSync(LoginAction{
		Username: config.Username,
		Secret:   config.Secret,
	})
	if err != nil {
		return nil, err
	}
	if res.Response() != "Success" {
		return nil, errors.New(res.GetString("Message"))
	}

	success = true
	return ami, nil
}

func (self *Command) Name() string {
	switch self.Type {
	case Action:
		return self.Action()
	case Response:
		return self.Response()
	case Event:
		return self.Event()
	default:
		return "UNKNOWN"
	}
}
func (self *Command) Response() string {
	return self.GetString("Response")
}

func (self *Command) Action() string {
	return self.GetString("Action")
}

func (self *Command) Event() string {
	return self.GetString("Event")
}
func (self *Command) Message() string {
	return self.GetString("Message")
}

func (self *Command) GetString(key string) string {
	return self.Headers[key].(string)
}

type ami struct {
	con           net.Conn
	scanner       *bufio.Scanner
	serverVersion string
	send          chan *_cmd
	recv          chan *Command
	cbs           map[int]chan CommandResponse
	aid           int
	subs          []chan<- *Command
	doneCh        chan bool
	done          bool
}

func newAMI() *ami {
	return &ami{
		recv:   make(chan *Command, 1024),
		send:   make(chan *_cmd, 1024),
		cbs:    make(map[int]chan CommandResponse),
		doneCh: make(chan bool, 1),
	}
}

func (self *ami) ServerVersion() string {
	return self.serverVersion
}

func (self *ami) WriteCommandSync(command interface{}) (*Command, error) {
	res := <-self.WriteCommand(command)
	return res.Response, res.Err
}

func (self *ami) WriteCommand(command interface{}) <-chan CommandResponse {
	cmd, err := buildAnyCommand(command)
	ch := make(chan CommandResponse, 1)
	if err != nil {
		ch <- CommandResponse{
			Err: err,
		}
	} else {
		c := &_cmd{
			response: ch,
			command:  cmd,
		}
		self.send <- c
	}
	log.Debugf("Queue Command %s", cmd.Name())
	return ch
}
func (self *ami) Close(err error) {
	if err != nil {
		log.WithError(err).Warnf("Close with error")
	}
	self.done = true
	self.doneCh <- true
	close(self.doneCh)
}
func (self *ami) fetch() {
	log.Debugf("AMI Fetch Start")
	scanner := self.scanner

	var headers map[string]interface{}
	for !self.done {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				self.Close(scanner.Err())
				break
			}
		}
		if headers == nil {
			headers = make(map[string]interface{})
		}
		line := scanner.Text()
		if len(line) == 0 {
			if len(headers) == 0 {
				continue
			}

			command, err := buildCommand(headers)
			if err != nil {
				log.WithError(err).WithField("headers", headers).Warnf("Build command failed")
				continue
			}
			log.WithField("headers", command.Headers).Debugf("Recv %v %s", command.Type, command.Name())
			self.recv <- command
			headers = nil
		} else {
			err := parseHeader(line, headers)
			if err != nil {
				log.WithError(err).WithField("header", line).Warnf("Parse header failed")
			}
		}
	}
	log.Debugf("AMI Fetch Stop")
}
func (self *ami) dispatch() {
	log.Debugf("AMI Dispatch Start")
	var c *_cmd
	ticker := time.Tick(time.Millisecond * 500)
	for {
		select {
		case c = <-self.send:
			headers := c.command.Headers
			self.aid++
			headers["ActionID"] = self.aid
			content := buildHeaders(headers)
			log.WithField("content", content).Debugf("Send")
			_, err := self.con.Write([]byte(content))
			if err != nil {
				c.response <- CommandResponse{
					Err: err,
				}
			} else {
				self.cbs[self.aid] = c.response
			}
		case r := <-self.recv:
			switch r.Type {
			case Event:
				for _, c := range self.subs {
					if len(c) < cap(c) {
						c <- r
					} else {
						log.Warnf("Subscribe channel fulled, will drop event")
					}
				}
			case Response:
				i, err := strconv.Atoi(fmt.Sprint(r.Headers["ActionID"]))
				if err != nil {
					log.WithField("ActionId", r.Headers["ActionID"]).Warnf("Failed to parse ActionId")
					continue
				}
				if r.Name() == "Error" {
					err = errors.New(r.Message())
				}
				if cb, ok := self.cbs[i]; ok {
					cb <- CommandResponse{
						Response: r,
						Err:      err,
					}
				} else {
					log.WithField("ActionId", i).Warnf("No callback found for action")
				}
			default:
				// Log
				log.WithField("command", r).Warnf("Invalid command")
			}
		case <-ticker:
		case <-self.doneCh:
			log.Debugf("AMI Dispatch Stop")
			return
		}
	}
}
func (self *ami) start() {
	if self.done {
		return
	}
	go self.fetch()
	go self.dispatch()
}

type _cmd struct {
	command  *Command
	response chan CommandResponse
}

func parseHeader(str string, headers map[string]interface{}) error {
	i := strings.IndexRune(str, ':')
	if i < 0 {
		return errors.New("Invalid header: " + str)
	}
	headers[str[0:i]] = strings.TrimSpace(str[i+1:])
	return nil
}

func buildHeaders(headers map[string]interface{}) string {
	buf := bytes.NewBufferString("")

	if v, ok := headers["Action"]; ok {
		buf.WriteString("Action")
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(v))
		buf.WriteString("\r\n")
	} else if v, ok := headers["Event"]; ok {
		buf.WriteString("Event")
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(v))
		buf.WriteString("\r\n")
	} else if v, ok := headers["Response"]; ok {
		buf.WriteString("Response")
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(v))
		buf.WriteString("\r\n")
	}

	for k, v := range headers {
		switch k {
		case "Action":
			continue
		case "Event":
			continue
		case "Response":
			continue
		}
		buf.WriteString(k)
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(v))
		buf.WriteString("\r\n")
	}

	buf.WriteString("\r\n")
	return buf.String()
}

var regTypeName = regexp.MustCompile("^(.*?)(Action|Event)$")

func buildAnyCommand(any interface{}) (*Command, error) {
	val := reflect.ValueOf(any)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	headers := make(map[string]interface{})

	match := regTypeName.FindStringSubmatch(val.Type().Name())
	if len(match) < 3 {
		return nil, errors.New("Invalid command type name: " + val.Type().Name())
	}
	headers[match[2]] = match[1]

	// NOTE Should I sort the key ?
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		headers[field.Name] = val.Field(i).Interface()
	}
	return buildCommand(headers)
}

func buildCommand(headers map[string]interface{}) (*Command, error) {
	command := &Command{
		Headers: headers,
	}
	if _, ok := headers["Action"]; ok {
		command.Type = Action
	} else if _, ok := headers["Response"]; ok {
		command.Type = Response
	} else if _, ok := headers["Event"]; ok {
		command.Type = Event
	} else {
		return nil, errors.New("Unknown command type")
	}
	return command, nil
}

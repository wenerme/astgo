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
	"github.com/wenerme/astgo/util"
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
	Action   string
	Response *Command
	Err      error
}

// A connection to Asterisk AMI
//
// This connection is NOT thread-safe
type Conn interface {
	ServerVersion() string
	WriteCommand(interface{}) <-chan CommandResponse
	WriteCommandSync(interface{}) (*Command, error)
	WriteCommandResponse(interface{}) interface{}
	Subscribe(...chan<- *Command) func()
	Close()

	SetDebug(bool)
}

type DialConf struct {
	Username  string
	Secret    string
	Listeners []chan<- *Command
}

func Connect(addr string) (c Conn, err error) {
	return connect(addr)
}

// Connect to server
func connect(addr string) (c *_con, err error) {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	var cn *_con
	defer func() {
		if err != nil {
			log.WithError(err).Warn("AMI Connect Failed")
			if cn != nil {
				cn.Close()
			} else {
				con.Close()
			}
		}
	}()

	scanner := bufio.NewScanner(con)
	cn = newCon()
	if scanner.Scan() {
		version := scanner.Text()
		match := regexp.MustCompile("Asterisk Call Manager/([0-9.]+)").FindStringSubmatch(version)
		if len(match) > 1 {
			cn.serverVersion = match[1]
		} else {
			err = errors.Errorf("Invalid server version: %s", version)
			return
		}
	}
	if scanner.Err() != nil {
		err = scanner.Err()
		return
	}
	cn.scanner = scanner
	cn.con = con
	go cn.start()

	return cn, nil
}

// Connect and Auth
func Dial(addr string, conf DialConf) (conn Conn, err error) {
	var c *_con
	c, err = connect(addr)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			c.Close()
		}
	}()

	c.Subscribe(conf.Listeners...)

	// Login
	if gr := c.WriteCommandResponse(LoginAction{
		Username: conf.Username,
		Secret:   conf.Secret,
	}).(*GeneralResponse); !gr.IsSuccess() {
		err = gr.ToError()
		return
	}

	return c, nil
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

// Response type
func (self *Command) Response() string {
	return self.GetString("Response")
}

// Action name
func (self *Command) Action() string {
	return self.GetString("Action")
}

// Event name
func (self *Command) Event() string {
	return self.GetString("Event")
}

// Message in response
func (self *Command) Message() string {
	return self.GetString("Message")
}

// Check is this a success response
func (self *Command) IsSuccess() bool {
	return self.Response() == "Success"
}

// Empty for not found
func (self *Command) GetString(key string) string {
	v, _ := self.Headers[key].(string)
	return v
}

// 0 for error or not found
func (self *Command) GetInt(key string) int {
	v, _ := strconv.Atoi(key)
	return v
}

type _con struct {
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
	dispatchAct   chan func() // Do something in dispatch goroutine
	debug         bool
}

func newCon() *_con {
	return &_con{
		recv:        make(chan *Command, 2048),
		send:        make(chan *_cmd, 1024),
		cbs:         make(map[int]chan CommandResponse),
		doneCh:      make(chan bool, 1),
		dispatchAct: make(chan func(), 2),
	}
}

func (self *_con) ServerVersion() string {
	return self.serverVersion
}

func (self *_con) Subscribe(listeners ...chan<- *Command) func() {
	if len(listeners) > 0 {
		self.dispatchAct <- func() {
			log.Debug("Append dial listeners")
			self.subs = append(self.subs, listeners...)
		}
	}
	return func() {
		// TODO
	}
}
func (self *_con) WriteCommandResponse(command interface{}) interface{} {
	res := <-self.WriteCommand(command)

	if hrt, ok := command.(HasResponseType); ok {
		rt := hrt.ResponseType()

		// Only GeneralResponse Can handle error response
		if _, ok := rt.(*GeneralResponse); !ok {
			if res.Err != nil {
				panic(res.Err)
			}
		}

		err := util.SetStruct(rt, res.Response.Headers)
		if err != nil {
			panic(res.Err)
		}
		return rt
	} else {
		if res.Err != nil {
			panic(res.Err)
		}
	}
	return nil
}
func (self *_con) WriteCommandSync(command interface{}) (*Command, error) {
	res := <-self.WriteCommand(command)
	return res.Response, res.Err
}

func (self *_con) WriteCommand(command interface{}) <-chan CommandResponse {
	cmd, err := buildAnyCommand(command)
	ch := make(chan CommandResponse, 1)
	if err != nil {
		ch <- CommandResponse{
			Action: cmd.Name(),
			Err:    err,
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
func (self *_con) SetDebug(debug bool) {
	self.debug = debug
}

func (self *_con) Close() {
	if err := self.con.Close(); err != nil {
		log.WithError(err).Warnf("Close connection return error")
	}
	self.done = true
	self.doneCh <- true
	close(self.doneCh)
}
func (self *_con) fetch() {
	log.Debugf("AMI Fetch Start")
	scanner := self.scanner

	var headers map[string]interface{}
	for !self.done {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				log.WithError(scanner.Err()).Warn("Fetch scanner failed")
				self.Close()
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
			if self.debug {
				log.WithField("headers", command.Headers).Debugf("Recv %v %s", command.Type, command.Name())
			}
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
func (self *_con) dispatch() {
	log.Debugf("AMI Dispatch Start")
	var c *_cmd
	ticker := time.Tick(time.Millisecond * 500)
	for {
		select {
		case a := <-self.dispatchAct:
			a()
		case c = <-self.send:
			headers := c.command.Headers
			self.aid++
			headers["ActionID"] = self.aid
			content := buildHeaders(headers)
			log.WithField("content", content).Debugf("Send")
			_, err := self.con.Write([]byte(content))
			if err != nil {
				c.response <- CommandResponse{
					Action: c.command.Action(),
					Err:    err,
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
func (self *_con) start() {
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

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

	log "github.com/sirupsen/logrus"
	"github.com/wenerme/astgo/util"
)

type CommandType int

const (
	Action CommandType = iota + 1
	Response
	Event
)

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
	Reconnect bool
	Debug     bool
	Listeners []chan<- *Command
}

func Connect(addr string) (c Conn, err error) {
	return connect(addr)
}

// Connect to server
func connect(addr string) (c *_con, err error) {
	con := newCon()
	con.addr = addr
	return con, con.connect()
}

func (self *_con) connect() (err error) {
	if self.done {
		return errors.New("Connection has bean closed")
	}

	con, err := net.DialTimeout("tcp", self.addr, time.Second*10)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			log.WithError(err).Warn("AMI Connect Failed")
			self.close()
		}
	}()

	scanner := bufio.NewScanner(con)
	if scanner.Scan() {
		version := scanner.Text()
		match := regexp.MustCompile("Asterisk Call Manager/([0-9.]+)").FindStringSubmatch(version)
		if len(match) > 1 {
			self.serverVersion = match[1]
		} else {
			err = errors.Errorf("Invalid server version: %s", version)
			return
		}
	}
	if scanner.Err() != nil {
		err = scanner.Err()
		return
	}
	self.scanner = scanner
	self.con = con
	go self.start()

	if self.username != "" {
		if res := self.WriteCommandResponse(LoginAction{
			Username: self.username,
			Secret:   self.secret,
		}).(*GeneralResponse); !res.IsSuccess() {
			err = res.ToError()
			return
		}
	}

	if self.autoReconnect {
		go self.reconnect()
	}

	return nil
}
func (self *_con) Ping() error {
	_, err := self.WriteCommandSync(PingAction{})
	return err
}
func (self *_con) reconnect() {
	if self.broken {
		log.Warn("AMI Reconnect routine already started")
		return
	}

	log.Debug("AMI Start reconnect routine")
	self.reconnection = true
	tick := time.Tick(10 * time.Second)
	for !self.done {
		if self.broken {
			if err := self.connect(); err != nil {
				log.WithError(err).Warn("Reconnect failed, try again later")
			} else {
				log.Info("Reconnect success")
				self.broken = false
				go self.fetch()
			}
		}

		select {
		case <-tick:
			if self.broken {
				continue
			}
			if err := self.Ping(); err != nil {
				log.WithError(err).Warn("Ping failed")
			}
		case <-self.doneCh:
			break
		case <-self.brokenCh:
			self.broken = true
			log.Debug("Detect broken")
			if err := self.con.Close(); err != nil {
				log.WithError(err).Warn("Close broken connection")
			}
			self.con = nil
		}
	}
	log.Debug("AMI Stop reconnect routine")
}

// Connect and Auth
func Dial(addr string, conf DialConf) (conn Conn, err error) {
	con := newCon()
	con.addr = addr
	con.username = conf.Username
	con.secret = conf.Secret
	con.autoReconnect = conf.Reconnect
	con.debug = conf.Debug

	con.Subscribe(conf.Listeners...)

	return con, con.connect()
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

// return nil if success
func (self *Command) ToError() error {
	if self.IsSuccess() {
		return nil
	}
	return errors.New(self.Message())
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
	autoReconnect bool
	reconnection  bool
	addr          string
	username      string
	secret        string
	con           net.Conn
	scanner       *bufio.Scanner
	serverVersion string
	send          chan *_request
	recv          chan *Command
	cbs           map[int]*_request
	aid           int
	subs          []chan<- *Command
	doneCh        chan bool
	broken        bool        // Connection broken
	brokenCh      chan bool   // Connection broken
	done          bool        // Closed
	dispatchAct   chan func() // Do something in dispatch goroutine
	debug         bool
}

func newCon() *_con {
	return &_con{
		recv:        make(chan *Command, 2048),
		send:        make(chan *_request, 1024),
		cbs:         make(map[int]*_request),
		doneCh:      make(chan bool, 1),
		brokenCh:    make(chan bool, 5),
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

		} else if res.Response == nil {
			// No response, error only
			r := &GeneralResponse{}
			r.Message = res.Err.Error()
			r.Response = "Error"
			return r
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
	ch := make(chan CommandResponse, 2)
	if err != nil {
		ch <- CommandResponse{
			Action: cmd.Name(),
			Err:    err,
		}
	} else {
		c := &_request{
			callback: ch,
			command:  cmd,
		}
		self.send <- c
	}
	if self.debug {
		log.Debugf("Queue Command %s", cmd.Name())
	}
	return ch
}
func (self *_con) SetDebug(debug bool) {
	self.debug = debug
}

func (self *_con) close() {
	if self.broken {
		// Ignore close
	} else {
		self.Close()
	}

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
				log.WithError(scanner.Err()).Warn("Scan failed")
				self.brokenConnection(scanner.Err())
				break
			}
		}
		if headers == nil {
			headers = make(map[string]interface{})
		}
		org := scanner.Text()
		line := strings.Trim(org, "\x00") // after reconnect, got a lot \x00
		if len(org) != len(line) {
			continue
		}
		if len(line) == 0 {
			if len(headers) == 0 {
				continue
			}

			command, err := buildCommand(headers)
			if err != nil {
				// After reconnect, will got a lot \x00, drop all
				log.WithError(err).WithField("headers", headers).Warnf("Build command failed")
				headers = nil
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
func (self *_con) brokenConnection(err error) {
	log.WithError(err).Warn("Connection broken")
	self.brokenCh <- true

	for i, r := range self.cbs {
		r.response(CommandResponse{
			Err: errors.New("Connection broken"),
		})
		delete(self.cbs, i)
	}
}
func (self *_con) dispatch() {
	log.Debugf("AMI Dispatch Start")
	var c *_request
	ticker := time.Tick(time.Second)
	for {
		select {
		case a := <-self.dispatchAct:
			a()
		case c = <-self.send:
			// Drop
			if self.con == nil {
				c.response(CommandResponse{
					Action: c.command.Action(),
					Err:    errors.New("Connection broken"),
				})
				log.WithField("command", c).Warn("Broken connection, Drop command")
				continue
			}
			headers := c.command.Headers
			self.aid++
			c.id = self.aid
			headers["ActionID"] = self.aid
			content := buildHeaders(headers)

			if self.debug {
				log.WithField("content", content).Debugf("Send")
			}

			self.con.SetWriteDeadline(time.Now().Add(5 * time.Second))
			_, err := self.con.Write([]byte(content))
			if err != nil {
				// Connection failed
				c.response(CommandResponse{
					Action: c.command.Action(),
					Err:    err,
				})
				log.WithError(err).Warn("Write command error")
				self.brokenConnection(err)
			} else {
				//log.WithField("len", n).Warn("Write command success")
				self.cbs[c.id] = c
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
				// 为这部分日志加上 command 信息
				// noinspection GoImportUsedAsName
				log := log.WithField("command", r)
				if err != nil {
					log.WithField("ActionId", r.Headers["ActionID"]).Warnf("Failed to parse ActionId")
					continue
				}
				err = r.ToError()
				if cb, ok := self.cbs[i]; ok {
					cb.response(CommandResponse{
						Response: r,
						Err:      err,
					})
					// FIXME 部分命令会响应两次, 例如 Originate Async
					delete(self.cbs, i)
				} else {
					log.WithField("ActionId", i).Warnf("No callback found for action")
				}
			default:
				// Log
				log.Warnf("Invalid command")
			}
		case <-ticker:
			deadline := time.Now().Add(-30 * time.Second)
			for i, r := range self.cbs {
				if r.time.Before(deadline) {
					r.response(CommandResponse{
						Err: errors.New("Timeout"),
					})
					log.WithField("action", r.command.Action()).WithField("delay", r.time.Sub(deadline).String()).WithField("id", r.id).Warnf("Timing out")
					delete(self.cbs, i)
				}
			}

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

type _request struct {
	id       int
	command  *Command
	callback chan CommandResponse
	time     time.Time
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
		// Drop null or empty
		if v == nil {
			continue
		}
		s := fmt.Sprint(v)
		if s == "" {
			continue
		}

		buf.WriteString(k)
		buf.WriteString(": ")
		buf.WriteString(s)
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

func (self *_request) response(response CommandResponse) {
	self.callback <- response
}

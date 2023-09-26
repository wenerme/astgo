package agi

import (
	"bufio"
	"context"
	"io"
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/wenerme/astgo/agi/agimodels"
	"go.uber.org/zap"
)

const (
	// StatusOK indicates operation was completed successfully.
	StatusOK = 200

	// StatusInvalid indicates invalid or unknown command.
	StatusInvalid = 510

	// StatusDeadChannel indicates the command cant be executed on a dead (closed, terminated, hung up) channel.
	StatusDeadChannel = 511

	// StatusEndUsage indicates end of proper usage, when the command returns its syntax.
	StatusEndUsage = 520
)

type Request struct {
}
type Response struct {
	Error        error  // Error received, if any
	Status       int    // HTTP-style status code received
	Result       int    // Result is the numerical return (if parseable)
	ResultString string // Result value as a string
	Value        string // Value is the (optional) string value returned
}

// Res returns the ResultString of a Response, as well as any error encountered.  Depending on the command, this is sometimes more useful than Val()
func (r *Response) Res() (string, error) {
	return r.ResultString, r.Error
}

// Err returns the error value from the response
func (r *Response) Err() error {
	return r.Error
}

// Val returns the response value and error
func (r *Response) Val() (string, error) {
	return r.Value, r.Error
}

type Session struct {
	Variables map[string]string
	r         io.Reader
	w         io.Writer
	ctx       context.Context
	conn      net.Conn
}

type HandlerFunc func(session *Session)

func NewSession(c context.Context, r io.Reader, w io.Writer, conn net.Conn) (*Session, error) {
	session := &Session{
		ctx:       c,
		r:         r,
		w:         w,
		conn:      conn,
		Variables: make(map[string]string),
	}
	return session, session.scanVariables()
}
func (a *Session) scanVariables() error {
	s := bufio.NewScanner(a.r)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		terms := strings.SplitN(s.Text(), ":", 2)
		if len(terms) == 2 {
			a.Variables[strings.TrimSpace(terms[0])] = strings.TrimSpace(terms[1])
		}
	}
	return s.Err()
}

// Close closes any network connection associated with the AGI instance
func (a *Session) Close() (err error) {
	if a.conn != nil {
		err = a.conn.Close()
		a.conn = nil
	}
	return
}
func ParseResponse(s string) *Response {
	sp := strings.SplitN(s, " ", 3)
	if len(sp) < 2 {
		return &Response{
			Error: errors.Errorf("invalid response: %q", s),
		}
	}
	r := &Response{}
	r.Status, r.Error = strconv.Atoi(sp[0])
	r.ResultString = sp[1][len("result="):]
	r.Result, _ = strconv.Atoi(r.ResultString)
	if len(sp) == 3 {
		r.Value = sp[2]
	}
	if r.Status != 200 {
		r.Error = errors.Errorf("status %v result %v extra %v", r.Status, r.ResultString, r.Value)
	}
	return r
}
func (a *Session) Command(cmd string) *Response {
	_, err := a.w.Write([]byte(cmd + "\n"))
	r := &Response{Error: err}
	if err != nil {
		return r
	}
	s := bufio.NewScanner(a.r)
	if s.Scan() {
		return ParseResponse(s.Text())
	}
	return r
}
func (a *Session) Client() *agimodels.Client {
	return &agimodels.Client{
		Handler: agimodels.HandlerFunc(func(cmd agimodels.Command) agimodels.Response {
			command, err := cmd.Command()
			if err != nil {
				return &Response{
					Error: err,
				}
			}
			return a.Command(command)
		}),
	}
}
func (a *Session) RequestVariable() *RequestVariable {
	rc := &RequestVariable{}
	rc.Load(a.Variables)
	return rc
}

// Listen binds an AGI HandlerFunc to the given TCP `host:port` address, creating a FastAGI service.
func Listen(addr string, handler HandlerFunc) error {
	if addr == "" {
		addr = "0.0.0.0:4573"
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "failed to bind server")
	}
	defer l.Close() // nolint: errcheck

	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.Wrap(err, "failed to accept TCP connection")
		}

		session, err := NewSession(nil, conn, conn, conn)
		if err != nil {
			return errors.Wrap(err, "failed init session")
		}
		go func() {
			defer session.Close()
			handler(session)
		}()
	}
}

type RequestVariable struct {
	Request      string
	Channel      string
	Language     string
	Type         string
	UniqueID     string
	Version      string
	CallerID     string
	CallerIDName string
	CallingPres  int // presentation of callerid
	CallingAni2  int
	CallingTon   int    // ast_channel_caller(chan)->id.number.plan
	CallingTns   int    // ast_channel_dialed(chan)->transit_network_select)
	DNID         string // ast_channel_dialed(chan)->number.str
	RDNIS        string // ast_channel_redirecting(chan)->from.number.valid, ast_channel_redirecting(chan)->from.number.str
	Context      string
	Extension    string
	Priority     int
	Enhanced     bool
	AccountCode  string
	ThreadID     int
	Args         []string
}

func (rc *RequestVariable) Load(m map[string]string) {
	var args []struct {
		i int
		v string
	}
	for k, v := range m {
		if !strings.HasPrefix(k, "agi_") {
			continue
		}
		name := k[4:]
		if strings.HasPrefix(name, "arg_") {
			i, err := strconv.Atoi(name[4:])
			if err != nil {
				//return errors.Wrapf(err, "invalid arg %v", name)
				zap.S().With("arg", name).Warn("skip invalid request arg")
			}
			args = append(args, struct {
				i int
				v string
			}{i: i, v: v})
			continue
		}
		if v == "unknown" {
			continue
		}
		var err error
		switch name {
		case "request":
			rc.Request = v
		case "channel":
			rc.Channel = v
		case "language":
			rc.Language = v
		case "type":
			rc.Type = v
		case "uniqueid":
			rc.UniqueID = v
		case "version":
			rc.Version = v
		case "callerid":
			rc.CallerID = v
		case "calleridname":
			rc.CallerIDName = v
		case "callingpres":
			rc.CallingPres, err = strconv.Atoi(v)
		case "callingani2":
			rc.CallingAni2, err = strconv.Atoi(v)
		case "callington":
			rc.CallingTon, err = strconv.Atoi(v)
		case "callingtns":
			rc.CallingTns, err = strconv.Atoi(v)
		case "dnid":
			rc.DNID = v
		case "rdnis":
			rc.RDNIS = v
		case "context":
			rc.Context = v
		case "extension":
			rc.Extension = v
		case "priority":
			rc.Priority, err = strconv.Atoi(v)
		case "enhanced":
			rc.Enhanced = v == "1.0"
		case "accountcode":
			rc.AccountCode = v
		case "threadid":
			rc.ThreadID, err = strconv.Atoi(v)
		}
		if err != nil {
			zap.S().With("err", err, "name", k, "value", v).Warn("failed to handle request variable")
		}
	}
	sort.Slice(args, func(i, j int) bool {
		return args[i].i < args[j].i
	})
	for _, v := range args {
		rc.Args = append(rc.Args, v.v)
	}
}

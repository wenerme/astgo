package agi

import (
	"bufio"
	"context"
	"github.com/pkg/errors"
	"github.com/wenerme/astgo/agi/agimodels"
	"io"
	"net"
	"strconv"
	"strings"
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

func NewSession(c context.Context, r io.Reader, w io.Writer) (*Session, error) {
	session := &Session{
		ctx:       c,
		r:         r,
		w:         w,
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
			command, err := cmd.Command(a.ctx)
			if err != nil {
				return &Response{
					Error: err,
				}
			}
			return a.Command(command)
		}),
	}
}

// Listen binds an AGI HandlerFunc to the given TCP `host:port` address, creating a FastAGI service.
func Listen(addr string, handler HandlerFunc) error {
	if addr == "" {
		addr = "localhost:4573"
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

		session, err := NewSession(nil, conn, conn)
		if err != nil {
			return errors.Wrap(err, "failed init session")
		}
		go handler(session)
	}
}

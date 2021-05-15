package agi

import (
	"bufio"
	"context"
	"github.com/wenerme/astgo/agi/agimodels"
	"io"
	"net"
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

	ctx  context.Context
	conn net.Conn
}

type HandlerFunc func(session *Session)

func NewSession(c context.Context, r io.Reader, w io.Writer) *Session {
	session := &Session{
		Variables: make(map[string]string),
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		terms := strings.SplitN(s.Text(), ":", 2)
		if len(terms) == 2 {
			session.Variables[strings.TrimSpace(terms[0])] = strings.TrimSpace(terms[1])
		}
	}
	return session
}

// Close closes any network connection associated with the AGI instance
func (a *Session) Close() (err error) {
	if a.conn != nil {
		err = a.conn.Close()
		a.conn = nil
	}
	return
}
func (a *Session) Command(cmd string) *Response {
	return nil
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

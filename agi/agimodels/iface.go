package agimodels

import (
	"context"
	"strconv"
	"strings"
)

type Command interface {
	Command() (string, error)
}
type needContext interface {
	SetContext(c context.Context)
}
type Response interface {
	Err() error
	Val() (string, error)
}

type Handler interface {
	Command(cmd Command) Response
}
type HandlerFunc func(cmd Command) Response

func (f HandlerFunc) Command(cmd Command) Response {
	return f(cmd)
}

type Client struct {
	Handler
}

func (c *Client) Answer() error {
	return c.Handler.Command(AnswerCommand{}).Err()
}
func (c *Client) Hangup() error {
	return c.Handler.Command(HangupCommand{}).Err()
}
func (c *Client) GetVariable(name string) (string, error) {
	return c.Handler.Command(GetVariableCommand{VariableName: name}).Val()
}
func (c *Client) SetVariable(name string, value string) (string, error) {
	return c.Handler.Command(SetVariableCommand{VariableName: name, Value: value}).Val()
}

func joinCommand(s []interface{}) string {
	sb := strings.Builder{}
	for _, param := range s {
		switch v := param.(type) {
		case string:
			sb.WriteString(v)
		case int:
			sb.WriteString(strconv.Itoa(v))
		case *string:
			if v == nil {
				goto DONE
			}
			sb.WriteString(*v)
		case *int:
			if v == nil {
				goto DONE
			}
			sb.WriteString(strconv.Itoa(*v))
		}
		sb.WriteRune(' ')
	}
DONE:
	return strings.TrimSpace(sb.String())
}

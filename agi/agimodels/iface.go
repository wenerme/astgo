package agimodels

import "context"

type Command interface {
	Command(c context.Context) (string, error)
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
	Handler Handler
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

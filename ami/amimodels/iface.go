package amimodels

import "errors"

type Event interface {
	EventTypeName() string
}
type Action interface {
	ActionTypeName() string
}
type HasActionID interface {
	GetActionID() string
	SetActionID(actionID string)
}
type Handler interface {
	Request(req *Request) *Response
}
type HandlerFunc func(req *Request) *Response

func (f HandlerFunc) Request(req *Request) *Response {
	return f(req)
}

type Client struct {
	Handler
}

func (cli *Client) Action(act Action, res interface{}, opts ...RequestOption) error {
	r, err := BuildRequest(act, res, opts...)
	if err != nil {
		return err
	}
	return cli.Request(r).Err()
}

type Response struct {
	Response  string // Success, Error
	ActionID  string
	Message   string
	Timestamp string // 1621260071.803488
	Headers   map[string]string
	Error     error
}

func (res *Response) Err() error {
	if res.Error != nil {
		return res.Error
	}
	if res.Response == "Error" {
		return errors.New(res.Message)
	}
	return nil
}

type Request struct {
	Action   Action
	Response *Response
}
type RequestOption func(r *Request) error

func BuildRequest(act Action, res interface{}, opts ...RequestOption) (r *Request, err error) {
	r = &Request{
		Action: act,
	}
	for _, v := range opts {
		if err = v(r); err != nil {
			return nil, err
		}
	}
	return
}

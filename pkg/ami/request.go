package ami

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

const attrActionID = "ActionID"

type requestOptions struct {
	Timeout    time.Duration
	OnComplete func(ctx context.Context, msg *Message, err error)
}
type RequestOption func(o *requestOptions) error

func RequestTimeout(d time.Duration) RequestOption {
	return func(o *requestOptions) error {
		o.Timeout = d
		return nil
	}
}

// RequestResponseCallback will case Conn.Request run async, will not timeout
func RequestResponseCallback(cb func(ctx context.Context, msg *Message, err error)) RequestOption {
	return func(o *requestOptions) error {
		o.OnComplete = cb
		return nil
	}
}

func (c *Conn) Request(r interface{}, opts ...RequestOption) (resp *Message, err error) {
	var msg *Message
	msg, err = ConvertToMessage(r)
	if err != nil {
		return nil, err
	}
	if msg.Type != MessageTypeAction {
		return nil, errors.Errorf("can only request action: %v", msg.Type)
	}

	async := &asyncMsg{
		id:     c.nextID(),
		msg:    msg,
		result: make(chan *asyncMsg, 1),
		ctx:    context.Background(),
	}
	o := requestOptions{
		Timeout: time.Second * 30,
	}
	for _, opt := range opts {
		if err = opt(&o); err != nil {
			return
		}
	}

	onComplete := o.OnComplete
	if onComplete != nil {
		async.cb = func(v *asyncMsg) {
			onComplete(v.ctx, v.resp, v.err)
		}
	}

	msg.SetAttr(attrActionID, async.id)

	if async.cb == nil {
		var cancel context.CancelFunc
		// todo allowed custom timeout
		async.ctx, cancel = context.WithTimeout(async.ctx, o.Timeout)
		defer cancel()
	}

	// enqueue
	c.pending <- async

	if async.cb != nil {
		return nil, errors.New("No response yet")
	}

	select {
	case <-async.ctx.Done():
		return nil, async.ctx.Err()
	case <-async.result:
		err = async.err
		if err == nil && async.resp != nil {
			err = async.resp.Error()
		}
		return async.resp, err
	}
}

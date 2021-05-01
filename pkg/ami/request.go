package ami

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

const attrActionID = "ActionID"

type RequestOption func(msg *asyncMsg) error

// RequestResponseCallback will case Conn.Request run async, will not timeout
func RequestResponseCallback(cb func(ctx context.Context, msg *Message, err error)) RequestOption {
	return func(msg *asyncMsg) error {
		msg.cb = func(v *asyncMsg) {
			cb(v.ctx, v.resp, v.err)
		}
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
	for _, opt := range opts {
		if err = opt(async); err != nil {
			return
		}
	}

	msg.SetAttr(attrActionID, async.id)

	if async.cb == nil {
		var cancel context.CancelFunc
		// allowed custom timeout
		async.ctx, cancel = context.WithTimeout(async.ctx, time.Second*30)
		defer cancel()
	}

	// enqueue
	c.pending <- async

	if async.cb != nil {
		return nil, errors.New("Async")
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

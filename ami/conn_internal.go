package ami

import (
	"context"
	"errors"
	"os"
	"time"
)

// promise like
type asyncMsg struct {
	id     string
	msg    *Message
	resp   *Message
	err    error
	result chan *asyncMsg
	cb     func(v *asyncMsg)
	ctx    context.Context
	done   bool
}

func (async *asyncMsg) complete() {
	if async.done {
		return
	}
	if async.result != nil {
		async.result <- async
		close(async.result)
	}
	if async.cb != nil {
		go func() {
			async.cb(async)
		}()
	}
	async.done = true
}

func (c *Conn) read(ctx context.Context) (err error) {
	log := c.logger
	log.Debug("start read loop")
	defer func() {
		log.Debug("done read loop")
	}()
	for {
		msg := &Message{}

		if err = msg.Read(c.reader); err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
			return
		}

		if msg.Type != "" {
			c.recv <- msg
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *Conn) loop(ctx context.Context) (err error) {
	ids := map[string]*asyncMsg{}
	defer func() {
		for _, v := range ids {
			v.err = err
			v.complete()
		}
	}()

	log := c.logger
	cleanTicker := time.NewTicker(5 * time.Second)

	c.logger.Debug("start event loop")
	defer func() {
		log.Debug("done event loop")
	}()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-cleanTicker.C:
			c.cleanUnsub()
		case async := <-c.pending:
			//err = c.conn.SetDeadline(time.Now().Add(1 * time.Second))
			//if err != nil {
			//	return
			//}

			// send pending message
			msg := async.msg
			//log.Sugar().With("id", async.id, "type", msg.Type, "name", msg.Name).Debug("send message")

			err = msg.Write(c.conn)

			if err != nil {
				async.err = err
				async.complete()
			} else if async.id != "" {
				ids[async.id] = async
			}
			c.onSend(async)
		case msg := <-c.recv:
			// received
			if msg.Type == MessageTypeResponse {
				id := msg.AttrString(attrActionID)
				if id != "" {
					async := ids[id]
					if async == nil {
						log.Sugar().With("id", id).Warn("response untracked")
					} else {
						async.resp = msg
						async.complete()
					}
				}
			}
			c.onRecv(msg)
		}
	}
}

func (c *Conn) onSend(msg *asyncMsg) {
	subs := c.subs

	for _, v := range subs {
		if v.unsub || !v.onSent {
			continue
		}
		ctx := v.ctx
		if ctx == nil {
			ctx = c.ctx
		}
		v.unsub = !v.f(ctx, msg.msg)
	}
}

func (c *Conn) onRecv(msg *Message) {
	subs := c.subs
	for _, v := range subs {
		if v.unsub || !v.onRecv {
			continue
		}
		ctx := v.ctx
		if ctx == nil {
			ctx = c.ctx
		}
		v.unsub = !v.f(ctx, msg)
	}
}

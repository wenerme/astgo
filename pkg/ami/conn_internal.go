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
}

func (async *asyncMsg) complete() {
	if async.result != nil {
		async.result <- async
		close(async.result)
	}
	if async.cb != nil {
		go func() {
			async.cb(async)
		}()
	}
}

func (c *Conn) read(ctx context.Context) (err error) {
	dl := 5 * time.Second
	log := c.logger
	log.Debug("start read loop")
	for {
		msg := &Message{}
		err = c.conn.SetReadDeadline(time.Now().Add(dl))
		if err != nil {
			return
		}
		// note partial read should not happen
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

func (c *Conn) loop(ctx context.Context) error {
	ids := map[string]*asyncMsg{}
	log := c.logger
	c.logger.Debug("start event loop")
	cleanTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-cleanTicker.C:
			c.cleanUnsub()
		case async := <-c.pending:
			// send pending message
			msg := async.msg
			err := msg.Write(c.conn)

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
	if !c.booted && msg.Name == "FullyBooted" {
		c.logger.Info("ami.Conn: FullyBooted")
		c.booted = true
		c.boot <- struct{}{}
		close(c.boot)
	}
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

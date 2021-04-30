package ami

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	amimodels "github.com/wenerme/astgo/pkg/ami/models"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

type Conn struct {
	g       *errgroup.Group
	conn    net.Conn
	reader  *bufio.Reader
	id      int64
	pending chan *asyncMsg
	recv    chan *Message
	ctx     context.Context
	closer  context.CancelFunc

	version string
	logger  *zap.Logger
	conf    *ConnectOptions
	subs    []*subscribe
	subLoc  sync.Mutex
}
type subscribe struct {
	f      SubscribeFunc
	ctx    context.Context
	unsub  bool
	onSent bool
	onRecv bool
}
type SubscribeFunc func(ctx context.Context, message *Message) bool
type ConnectOptions struct {
	Context        context.Context
	Timeout        time.Duration
	AllowReconnect bool
	Username       string // login username
	Secret         string // login secret
	Logger         *zap.Logger
	subscribers    []struct {
		sub  SubscribeFunc
		opts []SubscribeOption
	}
}
type ConnectOption func(c *ConnectOptions) error

func WithAuth(username string, secret string) ConnectOption {
	return func(c *ConnectOptions) error {
		c.Username = username
		c.Secret = secret
		return nil
	}
}

type SubscribeOption func(o *subscribe) error

func SubscribeSend() SubscribeOption {
	return func(o *subscribe) error {
		o.onSent = true
		return nil
	}
}

func SubscribeContext(ctx context.Context) SubscribeOption {
	return func(o *subscribe) error {
		o.ctx = ctx
		return nil
	}
}

func WithSubscribe(cb SubscribeFunc, opts ...SubscribeOption) ConnectOption {
	return func(c *ConnectOptions) error {
		c.subscribers = append(c.subscribers, struct {
			sub  SubscribeFunc
			opts []SubscribeOption
		}{sub: cb, opts: opts})
		return nil
	}
}
func Connect(addr string, opts ...ConnectOption) (conn *Conn, err error) {
	o := &ConnectOptions{
		Timeout: 10 * time.Second,
		Context: context.Background(),
		Logger:  zap.L(),
	}

	for _, v := range opts {
		if err = v(o); err != nil {
			return nil, err
		}
	}
	conf := *o

	c, err := net.DialTimeout("tcp", addr, conf.Timeout)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if e := c.Close(); e != nil {
				err = multierror.Append(err, e)
			}
		}
	}()

	conn = &Conn{
		conf:    &conf,
		logger:  o.Logger,
		recv:    make(chan *Message, 4096),
		pending: make(chan *asyncMsg, 100),
	}
	for _, sub := range o.subscribers {
		_, err = conn.Subscribe(sub.sub, sub.opts...)
		if err != nil {
			return nil, err
		}
	}
	return conn, conn.connect(c)
}

const attrActionID = "ActionID"

func (c *Conn) Request(ctx context.Context, msg *Message) (resp *Message, err error) {
	if ctx == nil {
		ctx = context.Background()
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()
	}

	if msg.Type != MessageTypeAction {
		return nil, errors.Errorf("can only request action: %v", msg.Type)
	}

	id := fmt.Sprint(atomic.AddInt64(&c.id, 1))
	msg.SetAttr(attrActionID, id)
	async := &asyncMsg{
		id:     id,
		msg:    msg,
		result: make(chan *asyncMsg, 1),
		ctx:    ctx,
	}

	c.pending <- async

	select {
	case <-async.ctx.Done():
		return nil, async.ctx.Err()
	case <-async.result:
		return async.resp, async.err
	}
}

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
		// todo partial read
		if err = msg.Read(c.reader); err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
			return
		}
		// log.Sugar().With("type", msg.Type, "name", msg.Name).Debug("read")

		if msg.Type != "" {
			c.recv <- msg
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func (c *Conn) Close() error {
	if c.closer != nil {
		c.closer()
		c.closer = nil
		return c.g.Wait()
	}
	return errors.New("closed or not init")
}

func (c *Conn) Subscribe(cb SubscribeFunc, opts ...SubscribeOption) (func(), error) {
	c.subLoc.Lock()
	defer c.subLoc.Unlock()

	sub := &subscribe{
		f:      cb,
		onRecv: true,
	}
	for _, v := range opts {
		err := v(sub)
		if err != nil {
			return nil, err
		}
	}

	c.subs = append(c.subs, sub)

	return func() {
		sub.unsub = true
	}, nil
}
func (c *Conn) clean() {
	c.subLoc.Lock()
	defer c.subLoc.Unlock()
	var neo []*subscribe
	subs := c.subs
	for _, sub := range subs {
		if !sub.unsub {
			neo = append(neo, sub)
		}
	}
	c.subs = neo
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
			c.clean()
		case async := <-c.pending:
			// send pending message
			msg := async.msg
			log.Sugar().With("type", msg.Type, "name", msg.Name, "attrs", msg.Attributes).Debug("sending")
			err := msg.Write(c.conn)
			log.Sugar().With("type", msg.Type, "name", msg.Name).Debug("send")

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
			log.Sugar().With("type", msg.Type, "name", msg.Name, "attrs", msg.Attributes).Debug("recv")
			c.onRecv(msg)
		}
	}
}

func (c *Conn) connect(conn net.Conn) (err error) {
	log := c.logger
	r := bufio.NewReader(conn)
	c.reader = r
	c.conn = conn

	line, err := r.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "scan ami initial line")
	}

	// check connection
	match := regexp.MustCompile("Asterisk Call Manager/([0-9.]+)").FindStringSubmatch(line)
	if len(match) > 1 {
		c.version = match[1]
		log.Sugar().With("version", c.version).Debug("AMI Version")
	} else {
		err = errors.Errorf("Invalid server header: %q", line)
		return
	}

	ctx := c.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	c.g, ctx = errgroup.WithContext(ctx)
	c.ctx = ctx
	c.g.Go(func() error {
		return c.read(ctx)
	})
	c.g.Go(func() error {
		return c.loop(ctx)
	})
	// manual close
	waitCtx, closer := context.WithCancel(ctx)
	c.g.Go(func() error {
		<-ctx.Done()
		closer()
		return conn.Close()
	})
	c.g.Go(func() error {
		<-waitCtx.Done()
		return errors.New("close")
	})
	c.closer = closer

	conf := c.conf
	if conf.Username != "" {
		var resp *Message
		resp, err = c.Request(nil, MustConvertToMessage(&amimodels.LoginAction{
			Username: conf.Username,
			Secret:   conf.Secret,
		}))

		if err != nil {
			err = errors.Wrap(err, "request login")
		} else if !resp.Success() {
			err = errors.Wrap(resp.Error(), "login")
		}
		if err != nil {
			return err
		}
		log.Info("login success")
	}
	return nil
}

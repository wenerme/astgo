package ami

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	amimodels "github.com/wenerme/astgo/ami/models"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net"
	"regexp"
	"sync/atomic"
	"time"
)

// CustomDialer can be used to specify any dialer, not necessarily
// a *net.Dialer.
type CustomDialer interface {
	Dial(network, address string) (net.Conn, error)
}

type ConnErrHandler func(*Conn, error)
type ConnHandler func(*Conn)

type ConnectOptions struct {
	Context        context.Context
	Timeout        time.Duration
	AllowReconnect bool
	Username       string // login username
	Secret         string // login secret
	Logger         *zap.Logger
	Dialer         CustomDialer
	OnConnectErr   ConnErrHandler
	OnConnected    ConnHandler
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
	opt := &ConnectOptions{
		Timeout: 10 * time.Second,
		Context: context.Background(),
		Logger:  zap.L(),
	}

	for _, v := range opts {
		if err = v(opt); err != nil {
			return nil, err
		}
	}
	if opt.Dialer == nil {
		opt.Dialer = &net.Dialer{
			Timeout: opt.Timeout,
		}
	}

	var id uint64
	conn = &Conn{
		ctx:     opt.Context,
		conf:    opt,
		logger:  opt.Logger,
		recv:    make(chan *Message, 4096),
		pending: make(chan *asyncMsg, 100),
		nextID: func() string {
			return fmt.Sprint(atomic.AddUint64(&id, 1))
		},
	}
	for _, sub := range opt.subscribers {
		_, err = conn.Subscribe(sub.sub, sub.opts...)
		if err != nil {
			return nil, err
		}
	}
	return conn, conn.dial(addr)
}

func (c *Conn) Close() error {
	if c.closer != nil {
		c.closed = true
		c.closer()
		c.closer = nil
		return c.g.Wait()
	}
	if c.closed {
		return nil
	}
	return errors.New("not init")
}

func (c *Conn) dial(addr string) (err error) {
	conf := c.conf

	if conf.AllowReconnect {
		// NOTE reconnect keep pending, but fail all async
		go func() {
			log := c.logger
			onErr := conf.OnConnectErr
			if onErr == nil {
				onErr = func(conn *Conn, err error) {
				}
			}

			var err error
			for !c.closed {
				err = c.dialOnce(addr)
				if err != nil {
					log.Sugar().With("err", err).Warn("ami.Conn: dial")

					onErr(c, err)
					// fixme improve wait strategy
					<-time.NewTimer(time.Second).C
					continue
				}
				if conf.OnConnected != nil {
					conf.OnConnected(c)
					log.Sugar().Info("ami.Conn: connected")
				}
				err = c.g.Wait()
				if err != nil {
					log.Sugar().With("err", err).Warn("ami.Conn: error")
					onErr(c, err)
				}
				c.g = nil
			}
			log.Sugar().Info("ami.Conn: stop reconnect, conn closed")
		}()
		return nil
	}
	return c.dialOnce(addr)
}

func (c *Conn) dialOnce(addr string) (err error) {
	conf := c.conf
	conn, err := conf.Dialer.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if e := conn.Close(); e != nil {
				err = multierror.Append(err, e)
			}
		}
	}()
	return c.connect(conn)
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
		resp, err = c.Request(amimodels.LoginAction{
			Username: conf.Username,
			Secret:   conf.Secret,
		}, RequestTimeout(2*time.Second))

		if err != nil {
			err = errors.Wrap(err, "request login")
		} else if !resp.Success() {
			err = errors.Wrap(resp.Error(), "login")
		}
		if err != nil {
			log.Sugar().With("err", err).Info("login failed")
			return err
		}
		log.Info("login success")
	}
	log.Sugar().Debug("do conn check ping")
	// be ready
	_, err = c.Request(amimodels.PingAction{})
	return
}

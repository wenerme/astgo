package ami

import (
	"bufio"
	"context"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net"
	"sync"
)

type Conn struct {
	g       *errgroup.Group
	conn    net.Conn
	reader  *bufio.Reader
	nextID  func() string
	pending chan *asyncMsg
	recv    chan *Message
	ctx     context.Context
	closer  context.CancelFunc
	closed  bool

	version string
	logger  *zap.Logger
	conf    *ConnectOptions
	subs    []*subscribe
	subLoc  sync.Mutex

	boot   chan struct{}
	booted bool
}

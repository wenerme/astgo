package ami

import (
	"bufio"
	"context"
	"net"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
}

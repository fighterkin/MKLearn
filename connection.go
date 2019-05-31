package mqbasic

import (
	"io"
)

type readDeadliner interface {
	SetReadDeadline(time.Time) error
}
type Connection struct {
	destructor sync.Once
	sendM      sync.Mutex
	m          sync.Mutex
	conn       io.ReadWriteCloser
	rpc        chan message
	writer     *writer
	sends      chan time.Time
	deadlines  chan readDeadLiner
	allocator  *allocator
	channels   map[uint16]*Channel
	noNotify   bool
	closes     []chan *Error
	blocks     []chan Blocking
	errors     chan *Error
	Major      int
	Minor      int
	Properties Table
	Locales    []string
	closed     int32
}

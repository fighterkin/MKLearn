package mqbasic

import (
	"reflect"
	"sync"
)

type Channel struct {
	destructor sync.Once
	m          sync.Mutex
	confirmM   sync.Mutex
	notify     sync.RWMutex
	connection *Connection
	rpc        chan message
	id         uint16
	closed     int32
	noNotify   bool
	closes     []chan *Error
	flows      []chan bool
	returns    []chan Return
	cancles    []chan string
	errors     chan *Error
	recv       func(*Channel, frame) error
	message    messageWithContet
	header     *headerFrame
	body       []byte
}

func newChannel(c *Connection, id uint64) Channel {

}

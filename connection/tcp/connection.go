package tcp

import (
	"errors"
	"net"
	"sync"
)

var (
	ErrCouldNotFoundConnection = errors.New("Could not found connection.")
)

type Receiver struct {
	conn     *net.TCPListener
	chRecv   chan recvObject
	chAccept chan acceptObject
	sync.Mutex
	clients []*net.TCPConn
}

func (c *Receiver) Close() {
	if c.conn == nil {
		return
	}
	c.Lock()
	defer c.Unlock()

	for _, client := range c.clients {
		client.Close()
	}

	close(c.chAccept)
	close(c.chRecv)
	c.conn.Close()
}
func (c *Receiver) Accept() chan acceptObject {
	go func() {
		for {
			conn, err := c.conn.AcceptTCP()
			if err != nil {
				c.chAccept <- acceptObject{nil, err}
				break
			}

			go func(conn *net.TCPConn) {
				c.Lock()
				defer c.Unlock()

				c.clients = append(c.clients, conn)

				c.chAccept <- acceptObject{conn.LocalAddr(), nil}
			}(conn)
		}
	}()
	return c.chAccept
}
func (c *Receiver) FindClient(addr net.Addr) *net.TCPConn {
	for _, client := range c.clients {
		if client.LocalAddr().String() != addr.String() {
			continue
		}
		return client
	}
	return nil
}
func (c *Receiver) Send(addr net.Addr, data []byte) (int, error) {
	for _, client := range c.clients {
		if addr != nil && client.LocalAddr().String() != addr.String() {
			continue
		}
		return client.Write(data)
	}
	return 0, ErrCouldNotFoundConnection
}

func (c *Receiver) Read(addr net.Addr) chan recvObject {
	conn := c.FindClient(addr)
	go func() {
		for {
			buf := make([]byte, 256)
			n, err := conn.Read(buf)
			if err != nil {
				c.chRecv <- recvObject{nil, conn.LocalAddr(), err}
			}
			c.chRecv <- recvObject{buf[:n], conn.LocalAddr(), nil}
		}
	}()
	return c.chRecv
}
func NewReceiver(addr string) (*Receiver, error) {
	var err error
	var laddr *net.TCPAddr

	laddr, err = net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return nil, err
	}

	return &Receiver{
		conn:     conn,
		chRecv:   make(chan recvObject, 10),
		chAccept: make(chan acceptObject, 10),
		clients:  make([]*net.TCPConn, 0, 8),
	}, nil
}

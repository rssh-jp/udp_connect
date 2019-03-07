package tcp

import (
	"net"
)

type Client struct {
	conn     *net.TCPConn
	chRecv   chan recvObject
	chAccept chan acceptObject
}

func (c *Client) Close() {
	if c.conn == nil {
		return
	}

	close(c.chAccept)
	close(c.chRecv)
	c.conn.Close()
}
func (c *Client) Send(data []byte) (int, error) {
	return c.conn.Write(data)
}

func (c *Client) Read(addr net.Addr) chan recvObject {
	go func() {
		for {
			buf := make([]byte, 256)
			n, err := c.conn.Read(buf)
			if err != nil {
				c.chRecv <- recvObject{nil, c.conn.LocalAddr(), err}
			}
			c.chRecv <- recvObject{buf[:n], c.conn.LocalAddr(), nil}
		}
	}()
	return c.chRecv
}
func NewClient(localAddr, remoteAddr string) (*Client, error) {
	var err error
	var laddr, raddr *net.TCPAddr

	if localAddr != "" {
		laddr, err = net.ResolveTCPAddr("tcp", localAddr)
		if err != nil {
			return nil, err
		}
	}

	raddr, err = net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", laddr, raddr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:     conn,
		chRecv:   make(chan recvObject, 10),
		chAccept: make(chan acceptObject, 10),
	}, nil
}

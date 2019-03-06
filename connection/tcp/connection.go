package tcp

import (
	"fmt"
	"net"
)

type recvObject struct {
	Data []byte
	Addr string
	Err  error
}
type Connect struct {
	conn   *net.TCPConn
	chRecv chan recvObject
}

func (c *Connect) Close() {
	close(c.chRecv)
	if c.conn == nil {
		return
	}
	c.conn.Close()
}
func (c *Connect) Read() chan recvObject {
	go func() {
		for {
			buf := make([]byte, 255)
			n, addr, err := c.conn.ReadFrom(buf)
			if err != nil {
				c.chRecv <- recvObject{nil, "", err}
				return
			}
			c.chRecv <- recvObject{buf[:n], addr.String(), nil}
		}
	}()
	return c.chRecv
}
func (c *Connect) Disconnect() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}
func (c *Connect) Send(data []byte) error {
	c.conn.Write(data)

	return nil
}
func (c *Connect) LocalAddr() string {
	return c.conn.LocalAddr().String()
}

func Create(localAddr, remoteAddr string) (*Connect, error) {
	var err error
	var laddr, raddr *net.UDPAddr

	if localAddr != "" {
		laddr, err = net.ResolveUDPAddr("udp", localAddr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if remoteAddr != "" {
		raddr, err = net.ResolveUDPAddr("udp", remoteAddr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("===", conn.RemoteAddr())
	fmt.Println("===", conn.LocalAddr())

	ret := Connect{
		conn:   conn,
		chRecv: make(chan recvObject, 10),
	}
	fmt.Println(conn)

	return &ret, nil
}

func CreateReceiver(localAddr string) (*Connect, error) {
	laddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ret := Connect{
		conn:   conn,
		chRecv: make(chan recvObject, 10),
	}

	return &ret, nil
}

package connection

import (
	"fmt"
	"net"
)

type Connect struct {
	conn *net.UDPConn
}

func (c *Connect) Close() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}
func (c *Connect) Read(cb func([]byte, error)) {
	go func() {
		for {
	        buf := make([]byte, 255)
			n, addr, err := c.conn.ReadFrom(buf)
			fmt.Println("addr : ", addr)
			if err != nil {
				cb(nil, err)
				return
			}
            fmt.Println("        1 connection.Read    : ", buf[:n])
			go func(data []byte) {
                fmt.Println("        2 connection.Read cb : ", data)
				cb(data, nil)
			}(buf[:n])
		}
	}()
	return
}
func (c *Connect) Disconnect() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}
func (c *Connect) Send(data []byte) error {
    fmt.Println("write word : ", data)
	c.conn.Write(data)

	return nil
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
		conn: conn,
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
		conn: conn,
	}
	fmt.Println(conn)

	return &ret, nil
}

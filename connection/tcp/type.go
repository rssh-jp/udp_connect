package tcp

import (
	"net"
)

type recvObject struct {
	Data []byte
	Addr net.Addr
	Err  error
}
type acceptObject struct {
	Addr net.Addr
	Err  error
}

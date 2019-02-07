package main

import (
	"fmt"

	"github.com/rssh-jp/udp_connect/connection"
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
)

func exec() {
	wk, err := connection.CreateReceiver(":5454")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wk.Disconnect()

	wk.Read(func(src []byte, addr string, err error) {
		if err != nil {
			wk.Disconnect()
			return
		}
		typ, obj, err := protocol.Deserialize(data.Deserialize(src, addr))
		if err != nil {
			wk.Disconnect()
			return
		}

		switch typ {
		case protocol.SysConnect:
			fmt.Println("connect")
		case protocol.SysAccessPoint:
			fmt.Println("ACCESS!!")
		case protocol.AppUser:
		case protocol.AppMessage:
		}
		fmt.Println(obj)
	})

	fmt.Println("Start")
	ch := make(chan struct{}, 1)

	<-ch
}
func main() {
	exec()
}

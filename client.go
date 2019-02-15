package main

import (
	"fmt"

	"github.com/rssh-jp/udp_connect/app"
	"github.com/rssh-jp/udp_connect/connection"
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
)

func receiver(localAddr string, chLocalAddr chan string, chClose chan error) {
	recv, err := connection.CreateReceiver(localAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer recv.Disconnect()

	chLocalAddr <- recv.LocalAddr()

	recv.Read(func(src []byte, addr string, err error) {
		if err != nil {
			recv.Disconnect()
			return
		}
		typ, obj, err := protocol.Deserialize(data.Deserialize(src, addr))
		if err != nil {
			recv.Disconnect()
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

	ch := make(chan struct{}, 1)
	<-ch
}

func main() {
	chLocalAddr := make(chan string, 1)
	chClose := make(chan error, 1)

	go receiver("", chLocalAddr, chClose)

	app.SendConnect(":5454")

	ch := make(chan struct{}, 1)
	<-ch
}

package main

import (
	"fmt"

	"github.com/rssh-jp/udp_connect/app"
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
	"github.com/rssh-jp/udp_connect/connection/udp"
)

func receiver(localAddr string, chLocalAddr chan string, chClose chan error) {
	recv, err := udp.CreateReceiver(localAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer recv.Disconnect()

	chLocalAddr <- recv.LocalAddr()

	chRecv := recv.Read()

	for {
		select {
		case res := <-chRecv:
			if res.Err != nil {
				recv.Disconnect()
				return
			}
			typ, obj, err := protocol.Deserialize(data.Deserialize(res.Data, res.Addr))
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
		}
	}
}

func main() {
	chLocalAddr := make(chan string, 1)
	chClose := make(chan error, 1)

	go receiver("", chLocalAddr, chClose)

	app.SendConnect(":5454")

	ch := make(chan struct{}, 1)
	<-ch
}

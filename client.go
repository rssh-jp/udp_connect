package main

import (
	"fmt"
	"log"

	"github.com/rssh-jp/udp_connect/app/udp"
	"github.com/rssh-jp/udp_connect/connection/tcp"
)

func main() {
	localAddr, err := udp.SendConnect(":5454")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(localAddr)

	conn, err := tcp.NewReceiver(localAddr)
	if err != nil {
		log.Fatal(err)
	}

	chAccept := conn.Accept()

	for {
		select {
		case accept := <-chAccept:
			chRecv := conn.Read(accept.Addr)
			for {
				select {
				case res := <-chRecv:
					fmt.Println("res : ", res)
				}
			}
		}
	}
}

package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
	"github.com/rssh-jp/udp_connect/connection/tcp"
	"github.com/rssh-jp/udp_connect/connection/udp"
)

func exec() {
	wk, err := udp.Create(":5454", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wk.Disconnect()

	var accessPoints sync.Map

	chRecv := wk.Read()

	fmt.Println("Start")

	for {
		select {
		case recv := <-chRecv:
			if recv.Err != nil {
				wk.Disconnect()
				return
			}
			fmt.Println("-----------------------", recv.Addr)
			_, obj, err := protocol.Deserialize(data.Deserialize(recv.Data, recv.Addr))
			if err != nil {
				wk.Disconnect()
				return
			}

			switch res := obj.(type) {
			case protocol.Connect:
				fmt.Println("connect", recv.Addr, res)
				if _, ok := accessPoints.Load(recv.Addr); ok {
					break
				}
				accessPoints.Store(recv.Addr, struct{}{})

				var addr string
				accessPoints.Range(func(key, value interface{}) bool {
					remote := key.(string)
					if strings.Contains(recv.Addr, remote) {
						return true
					}
					addr = remote
					return false
				})
				conn, err := tcp.NewClient("", recv.Addr)
				if err != nil {
					break
				}
				fmt.Println("addr : ", addr)
				conn.Send([]byte("aaaaaaa"))

			case protocol.AccessPoint:
				fmt.Println("ACCESS!!")
			case protocol.User:
			case protocol.Message:
			}
		}
	}
}
func main() {
	exec()
}

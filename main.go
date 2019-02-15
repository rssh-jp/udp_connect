package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rssh-jp/udp_connect/app"
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

	var accessPoints sync.Map

	wk.Read(func(src []byte, addr string, err error) {
		if err != nil {
			wk.Disconnect()
			return
		}
		fmt.Println("-----------------------", addr)
		_, obj, err := protocol.Deserialize(data.Deserialize(src, addr))
		if err != nil {
			wk.Disconnect()
			return
		}

		switch res := obj.(type) {
		case protocol.Connect:
			fmt.Println("connect")
			if _, ok := accessPoints.Load(res.Address); ok {
				break
			}
			accessPoints.Store(res.Address, struct{}{})

			fmt.Println(accessPoints)

			accessPoints.Range(func(key, value interface{}) bool {
				remote := key.(string)
				if strings.Contains(res.Address, remote) {
					return true
				}
				app.SendAccessPoint(res.Address, remote)
				fmt.Println("#####################", res.Address, remote)
				return false
			})

		case protocol.AccessPoint:
			fmt.Println("ACCESS!!")
		case protocol.User:
		case protocol.Message:
		}
		app.SendMessage(":5555", "aaaaaaaa")
	})

	fmt.Println("Start")

	ch := make(chan struct{}, 1)

	<-ch
}
func main() {
	exec()
}

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

	wk.Read(func(src []byte, err error) {
		if err != nil {
			wk.Disconnect()
			return
		}
		obj, err := protocol.Deserialize(data.Deserialize(src))
		fmt.Println(obj, err)
		if err != nil {
			wk.Disconnect()
			return
		}
		fmt.Println(obj)

		fmt.Println("-------------", string(src))
	})

	ch := make(chan struct{}, 1)

	<-ch
}
func main() {

	data, err := protocol.SerializeUser("nice")
	data2, err := protocol.Deserialize(data)
	//protocol.Deserialize([]byte(`{"protocol":1,"data":{"name":"aaaa","test":{"test2":10,"test3":"test3","test4":[11,12],"test5":[{"test6":13},{"test6":14}]}}}`))

	fmt.Println(data, err)
	fmt.Printf("%+v\n", data2)

	exec()
}

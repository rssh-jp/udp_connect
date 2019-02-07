package main

import (
	"fmt"

	"github.com/rssh-jp/udp_connect/app"
	"github.com/rssh-jp/udp_connect/connection"
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
)

func main() {
	conn, err := connection.Create("", ":5454")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("aaaaa : %d", i)
		app.SendMessage(conn, message)
	}

	app.SendUser(conn, "KKKKKKKKKKKKKKKKKKK")
	app.SendConnect(conn)
	app.SendAccessPoint(conn, "bbbbbbbbbbbbbbb")

	addr := conn.LocalAddr()

	conn.Close()

	app.SendMessage2(addr, ":5454", "new message")

	ch := make(chan struct{}, 1)
	<-ch
}

func sendMessage(conn *connection.Connect, message string) {
	sendData, err := protocol.SerializeMessage(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sendData))
	conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func sendUser(conn *connection.Connect, user string) {
	sendData, err := protocol.SerializeUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sendData))
	conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func sendConnect(conn *connection.Connect) {
	sendData, err := protocol.SerializeConnect()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sendData))
	conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func sendAccessPoint(conn *connection.Connect, accessPoint string) {
	sendData, err := protocol.SerializeAccessPoint(accessPoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sendData))
	conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

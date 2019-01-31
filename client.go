package main

import (
	"fmt"
	"net"

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

	//sendData, err := protocol.SerializeMessage("aaaaa")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//conn.Send(data.Serialize(sendData))

    //for i:=0; i<2; i++{
    //    message := fmt.Sprintf("aaaaa : %d", i)
    //    send(conn, message)

	//    //sendData, err := protocol.SerializeMessage(message)
	//    //if err != nil {
	//    //	fmt.Println(err)
	//    //	return
	//    //}

    //    //fmt.Println(string(sendData))
	//    //conn.Send(data.Serialize(sendData))
    //}
    send(conn, "aaaaaaaaaa")
    send(conn, "bbbbbbbbbb")

	ch := make(chan struct{}, 1)
	<-ch
}

func send(conn *connection.Connect, message string){
	sendData, err := protocol.SerializeMessage(message)
	if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Println(sendData)
    fmt.Println(string(sendData))
	conn.Send(data.Serialize(sendData))
}

func nice() {
	udpaddr, err := net.ResolveUDPAddr("udp", "localhost:5454")
	if err != nil {
		fmt.Println(err)
		return
	}

	//laddr, err := net.ResolveUDPAddr("udp", "localhost:5455")
	//if err != nil{
	//    fmt.Println(err)
	//    return
	//}

	conn, err := net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Create connection")

	conn.Write([]byte("data"))
	buf := make([]byte, 255, 255)

	// 最初の読み込みではアクセス先を取得
	//buf := make([]byte, 255)
	//n, err := conn.Read(buf)
	//if err != nil{
	//    fmt.Println(err)
	//    return
	//}

	for {
		n, err := conn.Read(buf)
		fmt.Println(n, err)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

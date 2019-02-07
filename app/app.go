package app

import (
	"github.com/rssh-jp/udp_connect/connection"
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
)

func SendMessage2(localAddr, remoteAddr string, message string) error {
	conn, err := connection.Create(localAddr, remoteAddr)
	if err != nil {
		return err
	}

	defer conn.Close()

	sendData, err := protocol.SerializeMessage(message)
	if err != nil {
		return err
	}

	return conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func SendMessage(conn *connection.Connect, message string) error {
	sendData, err := protocol.SerializeMessage(message)
	if err != nil {
		return err
	}

	return conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func SendUser(conn *connection.Connect, user string) error {
	sendData, err := protocol.SerializeUser(user)
	if err != nil {
		return err
	}

	return conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func SendConnect(conn *connection.Connect) error {
	sendData, err := protocol.SerializeConnect()
	if err != nil {
		return err
	}

	return conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func SendAccessPoint(conn *connection.Connect, accessPoint string) error {
	sendData, err := protocol.SerializeAccessPoint(accessPoint)
	if err != nil {
		return err
	}

	return conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

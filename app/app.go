package app

import (
	"github.com/rssh-jp/udp_connect/connection/data"
	"github.com/rssh-jp/udp_connect/connection/protocol"
	"github.com/rssh-jp/udp_connect/connection/udp"
)

func send(remoteAddr string, sendData []byte) (string, error) {
	conn, err := udp.Create("", remoteAddr)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	return conn.LocalAddr(), conn.Send(data.Serialize(sendData, conn.LocalAddr()))
}

func SendMessage(remoteAddr string, message string) (string, error) {
	sendData, err := protocol.SerializeMessage(message)
	if err != nil {
		return "", err
	}

	return send(remoteAddr, sendData)
}

func SendUser(remoteAddr string, user string) (string, error) {
	sendData, err := protocol.SerializeUser(user)
	if err != nil {
		return "", err
	}

	return send(remoteAddr, sendData)
}

func SendConnect(remoteAddr string) (string, error) {
	sendData, err := protocol.SerializeConnect()
	if err != nil {
		return "", err
	}

	return send(remoteAddr, sendData)
}

func SendAccessPoint(remoteAddr string, accessPoint string) (string, error) {
	sendData, err := protocol.SerializeAccessPoint(accessPoint)
	if err != nil {
		return "", err
	}

	return send(remoteAddr, sendData)
}

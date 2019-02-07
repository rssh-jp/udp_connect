package protocol

import (
	"encoding/json"

	"github.com/rssh-jp/udp_connect/convert"
)

func DeserializeConnect(data []byte) (Connect, error) {
	var ret Connect
	err := json.Unmarshal(data, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func SerializeConnect() ([]byte, error) {
	data := Connect{}
	return serialize(data)
}

func DeserializeAccessPoint(data []byte) (AccessPoint, error) {
	var ret AccessPoint
	err := json.Unmarshal(data, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func SerializeAccessPoint(address string) ([]byte, error) {
	data := AccessPoint{
		Address: address,
	}
	return serialize(data)
}

func DeserializeUser(data []byte) (User, error) {
	var ret User
	err := json.Unmarshal(data, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func SerializeUser(name string) ([]byte, error) {
	data := User{
		Name: name,
	}
	return serialize(data)
}

func DeserializeMessage(src []byte) (Message, error) {
	var ret Message
	err := json.Unmarshal(src, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func SerializeMessage(message string) ([]byte, error) {
	data := Message{
		Message: message,
	}
	return serialize(data)
}

func serialize(src interface{}) ([]byte, error) {
	var proto protocol
	switch src.(type) {
	case Connect:
		proto.Protocol = SysConnect
	case AccessPoint:
		proto.Protocol = SysAccessPoint
	case User:
		proto.Protocol = AppUser
	case Message:
		proto.Protocol = AppMessage
	}
	proto.Data = src

	dest, err := json.Marshal(proto)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func Deserialize(src []byte) (int, interface{}, error) {
	var d protocol
	err := json.Unmarshal(src, &d)
	if err != nil {
		return 0, nil, err
	}

	jsonStr := convert.MapInterface2string(d.Data)

	var retint int
	var retinterface interface{}
	var reterror error
	switch d.Protocol {
	case SysConnect:
		retinterface, reterror = DeserializeConnect([]byte(jsonStr))
		retint = SysConnect
	case SysAccessPoint:
		retinterface, reterror = DeserializeAccessPoint([]byte(jsonStr))
		retint = SysAccessPoint
	case AppUser:
		retinterface, reterror = DeserializeUser([]byte(jsonStr))
		retint = AppUser
	case AppMessage:
		retinterface, reterror = DeserializeMessage([]byte(jsonStr))
		retint = AppMessage
	}

	return retint, retinterface, reterror
}

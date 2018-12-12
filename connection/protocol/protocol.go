package protocol

import (
	"encoding/json"

	"github.com/rssh-jp/udp_connect/convert"
)

type protocol struct {
	Protocol int         `json:"protocol"`
	Data     interface{} `json:"data"`
}

func Deserialize(src []byte) (interface{}, error) {
	var d protocol
	err := json.Unmarshal(src, &d)
	if err != nil {
		return nil, err
	}
	jsonStr := convert.MapInterface2string(d.Data)
	switch d.Protocol {
	case AppUser:
		return DeserializeUser([]byte(jsonStr))
	case AppMessage:
		return DeserializeMessage([]byte(jsonStr))
	}

	return nil, nil
}

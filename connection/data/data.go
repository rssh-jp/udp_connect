package data

import (
	"github.com/rssh-jp/udp_connect/connection/data/crypto"
)

func Serialize(data []byte, key string) []byte {
	return crypto.Cryption(data, []byte(key))
}
func Deserialize(data []byte, key string) []byte {
	return crypto.Decryption(data, []byte(key))
}

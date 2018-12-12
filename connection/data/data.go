package data

import (
	"github.com/rssh-jp/udp_connect/connection/data/crypto"
)

var (
	key = []byte("テスト")
)

func Serialize(data []byte) []byte {
	crypto.Cryption(&data, key)
	return data
}
func Deserialize(data []byte) []byte {
	crypto.Decryption(&data, key)
	return data
}

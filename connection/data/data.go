package data

import (
	"github.com/rssh-jp/udp_connect/connection/data/crypto"
)

var (
	key = []byte("テスト")
)

func Serialize(data []byte) []byte {
	return crypto.Cryption(data, key)
}
func Deserialize(data []byte) []byte {
	return crypto.Decryption(data, key)
}

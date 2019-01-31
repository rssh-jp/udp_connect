package crypto

import (
	"math"
)

func addByte(a, b uint8) byte {
	work := a + b
	if work > math.MaxUint8 {
		return work - math.MaxUint8
	}
	return work
}

func Cryption(src, key []byte) []byte{
    ret := make([]byte, 0, len(src))
	for index, val := range src {
		keyIndex := index % len(key)
        ret = append(ret, addByte(val, key[keyIndex]))
	}
    return ret
}

func Decryption(src, key []byte) []byte{
    ret := make([]byte, 0, len(src))
	for index, val := range src {
		keyIndex := index % len(key)
        ret = append(ret, addByte(val, -key[keyIndex]))
	}
    return ret
}

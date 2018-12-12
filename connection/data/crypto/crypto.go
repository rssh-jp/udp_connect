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

func Cryption(src *[]byte, key []byte) {
	for index, val := range *src {
		keyIndex := index % len(key)
		(*src)[index] = addByte(val, key[keyIndex])
	}
}

func Decryption(src *[]byte, key []byte) {
	for index, val := range *src {
		keyIndex := index % len(key)
		(*src)[index] = addByte(val, -key[keyIndex])
	}
}

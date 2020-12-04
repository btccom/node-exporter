package peers

import (
	"crypto/rand"
	"encoding/binary"
)

func GetRandomUint64() (uint64, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	return binary.LittleEndian.Uint64(b), err
}

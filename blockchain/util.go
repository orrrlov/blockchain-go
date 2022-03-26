package blockchain

import (
	"bytes"
	"encoding/binary"

	"github.com/orrrlov/blockchain-go/errorz"
)

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		errorz.Panic(err)
	}

	return buff.Bytes()
}

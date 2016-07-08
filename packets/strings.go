package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const maxStringLength = 65535

func writeInt(i int) []byte {
	data := make([]byte, 2)

	binary.BigEndian.PutUint16(data, uint16(i))
	return data
}

func encode(s string) ([]byte, error) {
	if len(s) > maxStringLength {
		return nil, errors.New("String is too long")
	}

	buffer := bytes.Buffer{}
	buffer.Write(writeInt(len(s)))
	buffer.Write([]byte(s))

	return buffer.Bytes(), nil
}

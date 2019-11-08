package util

import (
	"bytes"
	"encoding/gob"
)

//GetBytestream converts the given value to a byteStream
func GetBytestream(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

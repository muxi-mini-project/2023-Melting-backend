package service

import (
	"encoding/base64"
)

func B64Encode(data string) []byte {
	ret := make([]byte, 0)
	base64.StdEncoding.Encode(ret, []byte(data))
	return ret
}
func B64Decode(data string) ([]byte, error) {
	ret, err := base64.StdEncoding.DecodeString(data)
	return ret, err
}

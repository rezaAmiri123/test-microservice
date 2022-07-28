package converter

import (
	"bytes"
	"encoding/json"
)

// Convert bytes to buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

//func BytesBufferToStruct(buf *bytes.Buffer) (interface{}, error) {
func BytesToStruct(byteData []byte) (interface{}, error) {
	reader := bytes.NewReader(byteData)
	var i interface{}
	err := json.NewDecoder(reader).Decode(i)
	return i, err
}

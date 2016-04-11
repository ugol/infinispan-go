package infinispan

import (
	"errors"
	"log"
)

// ResponseGet is the structure of a Get response
type ResponseGet struct {
	object []byte
}

func createGet(key []byte, messageID uint64, cacheName string) []byte {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, GetRequest, cacheName, false)
	p.EncodeBytes(key)
	return p.buf

}

// DecodeGetResponse creates a Get Response from a buffer
func (p *Buffer) DecodeGetResponse() (*ResponseGet, error) {

	var response = &ResponseGet{}
	if header, err := p.DecodeResponseHeader(); err == nil {

		response.object, _ = p.DecodeRawBytes()

		if header.opcode == ErrorResponse {
			log.Printf("%v", header)
			return response, errors.New(DecodeString(response.object))
		} else if header.opcode != GetResponse {
			log.Printf("%v", header)
			return response, errors.New("Not a possible GET Response")
		}

	} else {
		return response, err
	}

	return response, nil
}

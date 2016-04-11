package infinispan

import (
	"errors"
	"log"
)

// ResponsePut structure for Put Response
type ResponsePut struct {
	object []byte
}

func createPut(key []byte, value []byte, messageID uint64, cacheName string, lifespan string, maxidle string, previous bool) ([]byte, error) {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, PutRequest, cacheName, previous)
	p.EncodeBytes(key)
	err := p.AddLifespanAndMaxIdle(lifespan, maxidle)
	p.EncodeBytes(value)
	return p.buf, err

}

// DecodePutResponse creates a Put Response from a buffer
func (p *Buffer) DecodePutResponse() (*ResponsePut, error) {
	var response = &ResponsePut{}
	if header, err := p.DecodeResponseHeader(); err == nil {

		response.object, _ = p.DecodeRawBytes()
		if header.opcode == ErrorResponse {
			log.Printf("%v", header)
			return response, errors.New(DecodeString(response.object))
		} else if header.opcode != PutResponse {
			log.Printf("%v", header)
			return response, errors.New("Not a possible PUT Response")
		}
	} else {
		return response, err
	}

	return response, nil
}

package infinispan

import "errors"

// ResponsePut structure for Put Response
type ResponsePut struct {
	//empty at the moment
}

func createPut(key []byte, value []byte, messageID uint64, cacheName string) []byte {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, PutRequest, cacheName)
	p.EncodeBytes(key)
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeBytes(value)
	return p.buf

}

// DecodePutResponse creates a Put Response from a buffer
func (p *Buffer) DecodePutResponse() (*ResponsePut, error) {
	var response = &ResponsePut{}
	header, err := p.DecodeResponseHeader()

	if err == nil {
		if header.opcode != PutResponse {
			return response, errors.New("Not a PUT Response")
		}
	} else {
		return response, err
	}

	return response, nil
}

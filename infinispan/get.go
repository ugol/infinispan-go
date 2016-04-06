package infinispan

import "errors"

// ResponseGet is the structure of a Get response
type ResponseGet struct {
	object []byte
}

func createGet(key []byte, messageID uint64, cacheName string) []byte {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, GetRequest, cacheName)
	p.EncodeBytes(key)
	return p.buf

}

// DecodeGetResponse creates a Get Response from a buffer
func (p *Buffer) DecodeGetResponse() (*ResponseGet, error) {

	var response = &ResponseGet{}
	header, err := p.DecodeResponseHeader()

	if err == nil {
		if header.opcode != GetResponse {
			return response, errors.New("Not a GET Response")
		}
		response.object, _ = p.DecodeRawBytes()

	} else {
		return response, err
	}

	return response, nil
}

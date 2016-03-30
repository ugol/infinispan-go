package infinispan

import "errors"

// ResponseGet is the structure of a Get response
type ResponseGet struct {
	object []byte
}

func createGet(key []byte, messageID uint64, cachename string) []byte {

	p := NewBuffer([]byte{RequestMagic})
	p.EncodeVarint(messageID)
	p.EncodeRawBytes([]byte{Protocol20, GetRequest})
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeRawBytes([]byte{ClientIntelligenceBasic})
	p.EncodeRawBytes([]byte{0})
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

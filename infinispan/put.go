package infinispan

import "errors"

// ResponsePut structure for Put Response
type ResponsePut struct {
	//empty at the moment
}

func createPut(key []byte, value []byte, messageID uint64, cachename string) []byte {

	p := NewBuffer([]byte{RequestMagic})
	p.EncodeVarint(messageID)
	p.EncodeRawBytes([]byte{Protocol20, PutRequest})
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeRawBytes([]byte{ClientIntelligenceBasic})
	p.EncodeRawBytes([]byte{0})
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

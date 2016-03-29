package infinispan

import "errors"

type GetRes struct {
	messageID uint64
	status    byte
	topology  byte
	key       []byte
}

func createGet(key []byte, messageId uint64, cachename string) []byte {

	p := NewBuffer([]byte{RequestMagic})
	p.EncodeVarint(messageId)
	p.EncodeRawBytes([]byte{ProtocolVersionVersion20, GetRequest})
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeRawBytes([]byte{ClientIntelligenceBasic})
	p.EncodeRawBytes([]byte{0})
	p.EncodeBytes(key)
	return p.buf

}

func (p *Buffer) decodeMessageId() (uint64, error) {
	return p.DecodeVarint()
}

func (p *Buffer) decodeStatus() (byte, error) {
	return p.currentByte()
}

func (p *Buffer) decodeTopology() (byte, error) {
	return p.currentByte()
}

func (p *Buffer) decodeOpcode() (byte, error) {
	return p.currentByte()
}

func (p *Buffer) DecodeGetResponse() (*GetRes, error) {

	var response = &GetRes{}

	if err := p.decodeMagicResponse(); err == nil {
		response.messageID, _ = p.decodeMessageId()
		if op, _ := p.decodeOpcode(); op != GetResponse {
			return response, errors.New("Not a GET Response")
		}
		response.status, _ = p.decodeStatus()
		response.topology, _ = p.decodeTopology()
		response.key, _ = p.DecodeRawBytes()

	} else {
		return response, err
	}

	return response, nil
}

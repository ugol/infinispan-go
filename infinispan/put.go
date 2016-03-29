package infinispan

// PutRes structure fr Put Response
type PutRes struct {
	messageID uint64
	status    byte
	topology  byte
	key       []byte
}

func createPut(key []byte, value []byte, messageID uint64, cachename string) []byte {

	p := NewBuffer([]byte{RequestMagic})
	p.EncodeVarint(messageID)
	p.EncodeRawBytes([]byte{ProtocolVersionVersion20, PutRequest})
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeRawBytes([]byte{ClientIntelligenceBasic})
	p.EncodeRawBytes([]byte{0})
	p.EncodeBytes(key)
	p.EncodeRawBytes([]byte{0, 0})
	p.EncodeBytes(value)
	return p.buf

}

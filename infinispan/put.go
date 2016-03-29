package infinispan

type PutResponse struct {
	message_id uint64
	status byte
	topology byte
	key []byte
}

func createPut(key []byte, value []byte, messageId uint64, cachename string ) []byte {

	p := NewBuffer([]byte{REQUEST_MAGIC})
	p.EncodeVarint(messageId)
	p.EncodeRawBytes([]byte{PROTOCOL_VERSION_VERSION_20, PUT_REQUEST})
	p.EncodeRawBytes([]byte{0,0})
	p.EncodeRawBytes([]byte{CLIENT_INTELLIGENCE_BASIC})
	p.EncodeRawBytes([]byte{0})
	p.EncodeBytes(key)
	p.EncodeRawBytes([]byte{0,0})
	p.EncodeBytes(value)
	return p.buf

}

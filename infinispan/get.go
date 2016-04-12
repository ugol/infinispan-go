package infinispan

func createGet(key []byte, messageID uint64, cacheName string) []byte {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, GetRequest, cacheName, false)
	p.EncodeBytes(key)
	return p.buf

}

package infinispan

func createGet(key []byte, messageID uint64, cacheName string) []byte {
	return createReadOperation(GetRequest, key, messageID, cacheName)
}

func createRemove(key []byte, messageID uint64, cacheName string) []byte {
	return createReadOperation(RemoveRequest, key, messageID, cacheName)
}

func createContainsKey(key []byte, messageID uint64, cacheName string) []byte {
	return createReadOperation(ContainsKeyRequest, key, messageID, cacheName)
}

func createReadOperation(op byte, key []byte, messageID uint64, cacheName string) []byte {
	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, op, cacheName, false)
	p.EncodeBytes(key)
	return p.buf
}

package infinispan

func createPut(key []byte, value []byte, messageID uint64, cacheName string, lifespan string, maxidle string, previous bool) ([]byte, error) {

	p := NewBuffer([]byte{})
	p.CreateHeader(messageID, PutRequest, cacheName, previous)
	p.EncodeBytes(key)
	err := p.AddLifespanAndMaxIdle(lifespan, maxidle)
	p.EncodeBytes(value)
	return p.buf, err

}

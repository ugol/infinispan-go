package infinispan

// Buffer is a byte buffer which is the basis to decode/encode
type Buffer struct {
	buf   []byte
	index int
}

//NewBuffer creates a new Buffer
func NewBuffer(b []byte) *Buffer {
	return &Buffer{buf: b, index: 0}
}

//CreateHeader creates the basic header for every request
func (p *Buffer) CreateHeader(messageID uint64, opcode byte, cachename string, previous bool) {
	p.EncodeRawBytes([]byte{RequestMagic})
	p.EncodeVarint(messageID)
	p.EncodeRawBytes([]byte{Protocol25})
	p.EncodeRawBytes([]byte{opcode})
	p.EncodeString(cachename)
	if previous {
		p.EncodeRawBytes([]byte{1}) //Force return value flags
	} else {
		p.EncodeRawBytes([]byte{0}) //Empty client flags
	}
	p.EncodeRawBytes([]byte{ClientIntelligenceBasic})
	p.EncodeRawBytes([]byte{0}) //Client Topology ID
}

//AddLifespanAndMaxIdle encodes Lifespan and Maxidle
func (p *Buffer) AddLifespanAndMaxIdle(lifespan string, maxidle string) error {
	lD, lT, lE := parseDuration(lifespan)
	if lE != nil {
		return lE
	}
	mD, mT, mE := parseDuration(maxidle)
	if mE != nil {
		return mE
	}
	p.EncodeRawBytes([]byte{lT<<4 | mT})
	if lT < DefaultDuration {
		p.EncodeVarint(lD)
	}
	if mT < DefaultDuration {
		p.EncodeVarint(mD)
	}
	return nil
}

//DecodeResponseHeader decodes the common Response Header
func (p *Buffer) DecodeResponseHeader() (*ResponseHeader, error) {

	var response = &ResponseHeader{}

	if err := p.decodeMagicResponse(); err == nil {
		response.messageID, _ = p.decodeMessageID()
		response.opcode, _ = p.decodeOpcode()
		response.status, _ = p.decodeStatus()
		response.topology, _ = p.decodeTopology()
	} else {
		return response, err
	}

	return response, nil
}

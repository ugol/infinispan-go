package infinispan

//EncodeVarint encodes a vInt in the buffer
func (p *Buffer) EncodeVarint(x uint64) error {
	for x >= 1<<7 {
		p.buf = append(p.buf, uint8(x&0x7f|0x80))
		x >>= 7
	}
	p.buf = append(p.buf, uint8(x))
	return nil
}

//EncodeString encodes a string in the buffer
func (p *Buffer) EncodeString(s string) error {
	p.EncodeVarint(uint64(len(s)))
	p.buf = append(p.buf, s...)
	return nil
}

//EncodeRawBytes encodes a bytes slice in the buffer
func (p *Buffer) EncodeRawBytes(b []byte) error {
	p.buf = append(p.buf, b...)
	return nil
}

//EncodeBytes encodes a bytes slice in the buffer prepending the [] length
func (p *Buffer) EncodeBytes(b []byte) error {
	p.EncodeVarint(uint64(len(b)))
	p.buf = append(p.buf, b...)
	return nil
}

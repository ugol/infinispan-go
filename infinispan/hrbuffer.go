package infinispan

import (
	"errors"
	"fmt"
	"io"
)

// Buffer is a byte buffer which is the basis to decode/encode
type Buffer struct {
	buf   []byte
	index int
}

//NewBuffer creates a new Buffer
func NewBuffer(b []byte) *Buffer {
	return &Buffer{buf: b, index: 0}
}

//EncodeVarint encodes a vInt in the buffer
func (p *Buffer) EncodeVarint(x uint64) error {
	for x >= 1<<7 {
		p.buf = append(p.buf, uint8(x&0x7f|0x80))
		x >>= 7
	}
	p.buf = append(p.buf, uint8(x))
	return nil
}

//DecodeVarint decodes a vInt from the buffer
func (p *Buffer) DecodeVarint() (x uint64, err error) {
	// x, err already 0

	i := p.index
	l := len(p.buf)

	for shift := uint(0); shift < 64; shift += 7 {
		if i >= l {
			err = io.ErrUnexpectedEOF
			return
		}
		b := p.buf[i]
		i++
		x |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			p.index = i
			return
		}
	}

	// The number is too large to represent in a 64-bit value.
	err = errors.New("Integer overflow")
	return
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

//DecodeString decodesa protobuf encoded binary in a String
func DecodeString(b []byte) string {
	len := int(b[0])
	return string(b[1 : len+1])
}

//DecodeRawBytes gets an []bytes from the buffer
func (p *Buffer) DecodeRawBytes() (buf []byte, err error) {
	n, err := p.DecodeVarint()
	if err != nil {
		return nil, err
	}

	nb := int(n)
	if nb < 0 {
		return nil, fmt.Errorf("Bad byte length %d", nb)
	}
	end := p.index + nb
	if end < p.index || end > len(p.buf) {
		return nil, io.ErrUnexpectedEOF
	}

	buf = make([]byte, nb)
	copy(buf, p.buf[p.index:])
	p.index += nb
	return
}

func (p *Buffer) decodeMagicResponse() error {
	if p.buf[0] != ResponseMagic {
		return errors.New("Not a HotRod Response")
	}
	p.index++
	return nil
}

func (p *Buffer) currentByte() (byte, error) {
	i := p.index
	p.index++
	return p.buf[i], nil
}

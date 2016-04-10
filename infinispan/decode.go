package infinispan

import (
	"errors"
	"fmt"
	"io"
)

//ResponseHeader represents the Header for all responses
type ResponseHeader struct {
	messageID uint64
	opcode    byte
	status    byte
	topology  byte
}

func (h *ResponseHeader) String() string {
	return fmt.Sprintf("\nResponse Header\nMessage ID: %d\nOpcode: %#x (%s)\nStatus: %#x (%s)\nTopology: %d\n", h.messageID, h.opcode, responses[int(h.opcode)], h.status, status[int(h.status)], h.topology)
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

//DecodeString decodes a protobuf encoded binary in a String
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

func (p *Buffer) decodeMessageID() (uint64, error) {
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

func (p *Buffer) currentByte() (byte, error) {
	i := p.index
	p.index++
	return p.buf[i], nil
}

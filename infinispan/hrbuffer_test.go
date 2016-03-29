package infinispan

import (
	"testing"
	"fmt"
)

func TestEncodeInt(t *testing.T) {

	p := NewBuffer(nil)
	fmt.Printf("Original buffer: %v\n", p.buf)
	p.EncodeVarint(0)
	p.EncodeVarint(1)
	p.EncodeVarint(2)
	fmt.Printf("Original buffer: %v\n", p.buf)
	if p.buf[0] != byte(0) {
		t.Errorf("Wrong encoded value, expected %d, was %d", 0, p.buf[0])
	}
	if p.buf[1] != byte(1) {
		t.Errorf("Wrong encoded value, expected %d, was %d", 1, p.buf[1])
	}
	if p.buf[2] != byte(2) {
		t.Errorf("Wrong encoded value, expected %d, was %d", 2, p.buf[2])
	}

}

func TestEncodeString(t *testing.T) {

	original := "infinispan"
	plus := "rocks"

	p := NewBuffer(nil)
	fmt.Printf("Original buffer: %v\n", p.buf)
	p.EncodeString(original)
	fmt.Printf("First string: %v\n", p.buf)
	p.EncodeString(plus)

	if p.buf[0] != byte(len(original)) {
		expectedLen := int(p.buf[0])
		t.Errorf("Wrong length, expected %d, was %d", len(original), expectedLen)
	}

	encoded := string(p.buf[1:len(original) + 1])
	if (encoded != original) {
		t.Errorf("Wrong string, expected %s, was %s", original, encoded)
	}

	if p.buf[len(original) + 1] != byte(len(plus)) {
		expectedLen := int(p.buf[len(original) + 1])
		t.Errorf("Wrong length, expected %d, was %d", len(original) + 1, expectedLen)
	}

	encoded = string(p.buf[len(original) + 2:])
	if (encoded != plus) {
		t.Errorf("Wrong string, expected %s, was %s", plus, encoded)
	}

	fmt.Printf("Second string: %v\n", p.buf)

}

func TestEncodeEmptyString(t *testing.T) {

	p := NewBuffer(nil)
	original := ""
	p.EncodeString(original)

	if len(p.buf) != 1 {
		t.Errorf("Wrong length of empty string, expected %d, was %d", 1, len(p.buf))
	}

	if p.buf[0] != 0 {
		t.Errorf("Wrong empty string, byte value expected is %d, was %d", 0, int(p.buf[0]))
	}

}

func TestDecodeString(t *testing.T) {

	buf := []byte{
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	ugol := DecodeString(buf)

	if (ugol != "ugol") {
		t.Errorf("Wrong string, expected %s, was %s", "ugol", ugol)
	}

}


func TestDecodeEmptyString(t *testing.T) {

	buf := []byte{
		0x00, //Value length
	}

	empty := DecodeString(buf)

	if (empty != "") {
		t.Errorf("Wrong string, expected %s, was %s", "", empty)
	}

}



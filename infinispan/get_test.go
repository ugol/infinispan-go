package infinispan

import (
	"bytes"
	"testing"
)

func TestCreateGet(t *testing.T) {

	expected_get := []byte{
		0xA0, //MAGIC BYTE
		0x00, //Message Id
		0x14, //Protocol version
		0x03, //Operation
		0x00, //Cache name length ("" = default)
		0x00, //Flags
		0x01, //Client intelligence
		0x00, //Client Topology ID
		0x01, //Key length
		0x31, //Key: 1
	}

	MakeId(0)
	get := createGet([]byte("1"), <-id, "")

	if !bytes.Equal(expected_get, get) {
		t.Errorf("Expected %v, was %v", expected_get, get)
	}

}

package infinispan

import (
	"bytes"
	"testing"
)

func TestCreatePut(t *testing.T) {

	expectedPut := []byte{
		0xA0, //MAGIC BYTE
		0x00, //Message Id
		0x14, //Protocol version
		0x01, //Operation
		0x00, //Cache name length ("" = default)
		0x00, //Flags
		0x01, //Client intelligence
		0x00, //Client Topology ID
		0x01, //Key length
		0x32, //Key: 2
		0x00, //Lifespan
		0x00, //Max idle
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	MakeId(0)
	put := createPut([]byte("2"), []byte("ugol"), <-id, "")

	if !bytes.Equal(expectedPut, put) {
		t.Errorf("Expected %v, was %v", expectedPut, put)
	}

}

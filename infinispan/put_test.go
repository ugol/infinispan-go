package infinispan

import (
	"bytes"
	"testing"
)

func TestCreatePut(t *testing.T) {

	expectedPut := []byte{
		0xA0, //MAGIC BYTE
		0x00, //Message Id
		0x19, //Protocol version
		0x01, //Operation
		0x00, //Cache name length ("" = default)
		0x00, //Flags
		0x01, //Client intelligence
		0x00, //Client Topology ID
		0x01, //Key length
		0x32, //Key: 2
		0x77, //Default Lifespan & Max idle
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	MakeID(0)
	put, _ := createPut([]byte("2"), []byte("ugol"), <-id, "", "0", "0")

	if !bytes.Equal(expectedPut, put) {
		t.Errorf("Expected %v, was %v", expectedPut, put)
	}

}

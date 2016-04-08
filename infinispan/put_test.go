package infinispan

import (
	"bytes"
	"testing"
)

func TestCreateDefaultPut(t *testing.T) {

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
		0x77, //Default Lifespan & Maxidle
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

func TestCreatePutWithLifespan(t *testing.T) {

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
		0x17, //ms Lifespan & Default Maxidle
		0x05, //Duration in ms
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	MakeID(0)
	put, _ := createPut([]byte("2"), []byte("ugol"), <-id, "", "5ms", "0")

	if !bytes.Equal(expectedPut, put) {
		t.Errorf("Expected %v, was %v", expectedPut, put)
	}

}

func TestCreatePutWithMaxidle(t *testing.T) {

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
		0x71, //Default Lifespan & ms Maxidle
		0x05, //Maxidle in ms
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	MakeID(0)
	put, _ := createPut([]byte("2"), []byte("ugol"), <-id, "", "0", "5ms")

	if !bytes.Equal(expectedPut, put) {
		t.Errorf("Expected %v, was %v", expectedPut, put)
	}

}

func TestCreatePutWithLifespanAndMaxidle(t *testing.T) {

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
		0x12, //ms Lifespan & ns Maxidle
		0x05, //Lifespan in ms
		0x04, //Maxidle in ms
		0x04, //Value length
		0x75, //u
		0x67, //g
		0x6F, //o
		0x6c, //l
	}

	MakeID(0)
	put, _ := createPut([]byte("2"), []byte("ugol"), <-id, "", "5ms", "4ns")

	if !bytes.Equal(expectedPut, put) {
		t.Errorf("Expected %v, was %v", expectedPut, put)
	}

}

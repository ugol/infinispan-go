package infinispan

import (
	"bytes"
	"testing"
)

func TestSimplePutAndGet(t *testing.T) {

	c, err := NewConnection("127.0.0.1:11222")
	if err != nil {
		t.Error(err.Error())
		return
	}
	c.Put([]byte("1"), []byte("foo"))
	c.Put([]byte("2"), []byte("bar"))
	c.Put([]byte("3"), []byte("ugol"))
	ugol, err := c.Get([]byte("3"))
	if err != nil {
		t.Error(err.Error())
	}

	if !bytes.Equal([]byte("ugol"), ugol.object) {
		t.Errorf("Expected %v, was %v", []byte("ugol"), ugol)
	}
	c.Close()

}

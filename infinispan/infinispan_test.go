package infinispan

import (
	"bytes"
	"testing"
)

const conf = `
		{
		 "servers":[
			 {"host": "127.0.0.1", "port": 11222}
		 ],
		 "cacheName": ""
		}
	`

func TestSimplePutAndGet(t *testing.T) {

	if c, err := NewClientJSON(conf); err == nil {
		defer c.Close()

		c.Put([]byte("1"), []byte("foo"))
		c.Put([]byte("2"), []byte("bar"))
		c.Put([]byte("3"), []byte("ugol"))
		if ugol, err1 := c.Get([]byte("3")); err1 == nil {
			if !bytes.Equal([]byte("ugol"), ugol.object) {
				t.Errorf("Expected %v, was %v", []byte("ugol"), ugol)
			}
		} else {
			t.Error(err1.Error())
		}
	} else {
		t.Error(err.Error())
		return
	}

}

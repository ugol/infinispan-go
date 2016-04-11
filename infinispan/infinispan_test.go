package infinispan

import (
	"bytes"
	"testing"
	"time"
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

		if _, errPut := c.Put([]byte("1"), []byte("foo")); errPut != nil {
			t.Error(errPut.Error())
		}

		if _, errPut := c.Put([]byte("2"), []byte("bar")); errPut != nil {
			t.Error(errPut.Error())
		}

		if _, errPut := c.Put([]byte("3"), []byte("ugol")); errPut != nil {
			t.Error(errPut.Error())
		}

		if ugol, errGet := c.Get([]byte("3")); errGet == nil {
			if !bytes.Equal([]byte("ugol"), ugol.object) {
				t.Errorf("Expected %v, was %v", []byte("ugol"), ugol)
			}
		} else {
			t.Error(errGet.Error())
		}

		if notFound, errGet := c.Get([]byte("4")); errGet == nil {
			if !bytes.Equal([]byte(""), notFound.object) {
				t.Errorf("Expected %v, was %v", []byte(""), notFound)
			}
		} else {
			t.Error(errGet.Error())
		}

	} else {
		t.Error(err.Error())
		return
	}

}

func TestLifespanPut(t *testing.T) {

	if c, err := NewClientJSON(conf); err == nil {
		defer c.Close()

		opts := map[string]string{
			"lifespan": "10ms",
		}

		c.PutWithOptions([]byte("100"), []byte("T"), opts)

		if found, errGet := c.Get([]byte("100")); errGet == nil {
			if !bytes.Equal([]byte("T"), found.object) {
				t.Errorf("Expected %v, was %v", []byte("T"), found)
			}
		} else {
			t.Error(errGet.Error())
		}

		time.Sleep(20 * time.Millisecond)

		if notFound, errGet := c.Get([]byte("100")); errGet == nil {
			if !bytes.Equal([]byte(""), notFound.object) {
				t.Errorf("Expected %v, was %v", []byte(""), notFound)
			}
		} else {
			t.Error(errGet.Error())
		}

	} else {
		t.Error(err.Error())
		return
	}

}

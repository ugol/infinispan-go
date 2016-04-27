package infinispan

import (
	"math"
	"testing"
)

func TestSimpleIds(t *testing.T) {

	id0 := make(chan uint64)
	closeCh := make(chan bool)
	Start(0, id0, closeCh)
	x := <-id0
	x = <-id0
	x = <-id0
	x = <-id0
	if x != 3 {
		t.Errorf("Wrong id, expected %d, was %d", 3, x)
	}
	closeCh <- true
}

func TestMaxId(t *testing.T) {

	idMax := make(chan uint64)
	closeCh := make(chan bool)
	Start(math.MaxUint64, idMax, closeCh)
	x := <-idMax
	if x != math.MaxUint64 {
		t.Errorf("Wrong id, expected max uint64 %d, was %d", uint64(math.MaxUint64), x)
	}
	x = <-idMax
	if x != 0 {
		t.Errorf("Wrong id, should start again from %d, but was %d", 0, x)
	}
	closeCh <- true

}

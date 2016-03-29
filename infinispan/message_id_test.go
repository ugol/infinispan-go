package infinispan

import (
	"math"
	"testing"
)

func TestSimpleIds(t *testing.T) {

	MakeID(0)
	x := <-id
	x = <-id
	x = <-id
	x = <-id
	if x != 3 {
		t.Errorf("Wrong id, expected %d, was %d", 3, x)
	}

}

func TestMaxId(t *testing.T) {

	MakeID(math.MaxUint64)
	x := <-id
	if x != math.MaxUint64 {
		t.Errorf("Wrong id, expected max uint64 %d, was %d", uint64(math.MaxUint64), x)
	}
	x = <-id
	if x != 0 {
		t.Errorf("Wrong id, should start again from %d, but was %d", 0, x)
	}
}

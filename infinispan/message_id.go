package infinispan

var id chan uint64

func init() {
	id = MakeID(0)
}

// MakeID explicitly creates the IDs channel, for example if you don't want to start from 0
func MakeID(start uint64) chan uint64 {

	ch := make(chan uint64)
	go func() {
		for i := uint64(start); ; i++ {
			ch <- i
		}
	}()
	return ch
}

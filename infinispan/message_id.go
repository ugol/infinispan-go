package infinispan

var id chan uint64

func init() {
	MakeID(0)
}

// MakeID explicitly creates the IDs channel, for example if you don't want to start from 0
func MakeID(start uint64) {

	id = make(chan uint64)
	go func() {
		for i := uint64(start); ; i++ {
			id <- i
		}
	}()

}

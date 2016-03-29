package infinispan

var id chan uint64

func init() {
	MakeId(0)
}

func MakeId(start uint64) {

	id = make(chan uint64)
	go func() {
		for i:=uint64(start);; i++ {
			id <- i
		}
	}()

}
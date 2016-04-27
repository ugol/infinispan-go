package infinispan

// Start explicitly creates the IDs channel, useful in testing
func Start(start uint64, ch chan uint64, done chan bool) {

	go func() {

		i := uint64(start)
		for {
			select {
			case <-done:
				close(ch)
				return
			case ch <- i:
				i++
			}

		}
	}()

}

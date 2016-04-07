package infinispan

import "testing"

func TestDurationInMs(t *testing.T) {
	duration := "5000 ms"
	d, time, err := parseDuration(duration)

	if err != nil {
		t.Errorf("Shouldn't generate an error, but it generated '%s'", err)
	}

	if d != 5000 {
		t.Errorf("Expected %d, was %d", 5000, d)
	}

	if time != 2 {
		t.Errorf("'ms' should be encoded with byte %d, but was %d", 1, time)
	}
}

func TestSplit(t *testing.T) {

	/*
		splitter = regexp.MustCompile(`^(\d+)\s{0,}(s|ms|ns|μs|m|h|d)\s{0,}$`)
		result2 := splitter.FindStringSubmatch("5000ms")
		for k, v := range result2 {
			fmt.Printf("%d. %s\n", k, v)
		}
	*/

}

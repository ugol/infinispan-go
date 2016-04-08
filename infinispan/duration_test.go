package infinispan

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParsingDuration(t *testing.T) {

	var tests = []struct {
		input    string
		duration uint64
		time     byte
		err      error
	}{
		{"5000 ms", 5000, 1, nil},
		{"5000ms", 5000, 1, nil},
		{"20s", 20, 0, nil},
		{"100 ns", 100, 2, nil},
		{"100000μs", 100000, 3, nil},
		{"2m", 2, 4, nil},
		{"1h", 1, 5, nil},
		{"1d", 1, 6, nil},
		{"-125", 0, InfiniteDuration, nil},
		{"0", 0, DefaultDuration, nil},
		{"5000", 0, DefaultDuration, fmt.Errorf("Positive duration 5000 provided without time unit")},
		{"1mm", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1mm")},
		{"1hh", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1hh")},
		{"1dd", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1dd")},
		{"1dd", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1dd")},
		{"1hd", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1hd")},
		{"1 sm", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1 sm")},
		{"1sm", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1sm")},
		{"1n", 0, DefaultDuration, fmt.Errorf("Unknown duration format for 1n")},
		{"-1s", 0, DefaultDuration, fmt.Errorf("Unknown duration format for -1s")},
	}

	for _, test := range tests {
		d, time, err := parseDuration(test.input)

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("Error should be '%v', but it was '%v'", test.err, err)
		}

		if d != test.duration {
			t.Errorf("Expected %d, was %d", test.duration, d)
		}

		if time != test.time {
			t.Errorf("With '%s' expected duration in byte is %d, but was %d", test.input, test.time, time)
		}

	}

}

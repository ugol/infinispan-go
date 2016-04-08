package infinispan

import (
	"fmt"
	"regexp"
	"strconv"
)

var durationRE = regexp.MustCompile(`^(\d+)\s{0,}(s|ms|ns|μs|m|h|d)\s{0,}$`)

func timeUnitToByte(unit string) byte {
	switch unit {
	case "s":
		return 0
	case "ms":
		return 1
	case "ns":
		return 2
	case "μs":
		return 3
	case "m":
		return 4
	case "h":
		return 5
	case "d":
		return 6
	default:
		panic("Can't be here, any valid time unit should be checked by the regex")
	}
}

func parseDuration(duration string) (uint64, byte, error) {

	if d, err := strconv.Atoi(duration); err == nil {
		if d < 0 {
			return 0, InfiniteDuration, nil
		}
		if d == 0 {
			return 0, DefaultDuration, nil
		}
		return 0, DefaultDuration, fmt.Errorf("Positive duration %d provided without time unit", d)
	}

	parsedDuration := durationRE.FindStringSubmatch(duration)
	if parsedDuration == nil || len(parsedDuration) != 3 {
		return 0, DefaultDuration, fmt.Errorf("Unknown duration format for %s", duration)
	}

	d, _ := strconv.Atoi(parsedDuration[1])
	t := timeUnitToByte(parsedDuration[2])
	return uint64(d), t, nil

}

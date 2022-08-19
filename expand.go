// Parsing of numeric ranges from string
package expandrange

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Range type
type Range []uint64

// Checking if a number is in a range
func (rng *Range) InRange(num uint64) bool {
	for _, i := range *rng {
		if num == i {
			return true
		}
	}
	return false
}

func (rng *Range) append(u uint64) {
	if !rng.InRange(u) {
		*rng = append(*rng, u)
	}
}

// Parse string in range
func (rng *Range) Parse(s string) error {
	if rng.Len() > 0 {
		rng = &Range{}
	}
	if s == "" {
		return nil
	}
	for _, i := range strings.Split(s, ",") {
		if strings.Contains(i, "-") {
			r := strings.Split(i, "-")
			if len(r) != 2 {
				return fmt.Errorf("the range '%s' contains more than two values", i)
			}
			min, err := strconv.ParseUint(r[0], 10, 32)
			if err != nil {
				return err
			}
			max, err := strconv.ParseUint(r[1], 10, 32)
			if err != nil {
				return err
			}
			if max < min {
				return fmt.Errorf("%d must be greater than %d", max, min)
			}
			for i := min; i < max+1; i++ {
				rng.append(i)
			}
		} else {
			u, err := strconv.ParseUint(i, 10, 32)
			if err != nil {
				return err
			}
			rng.append(u)
		}
	}
	return nil
}

// Len is length of range
func (rng Range) Len() int {
	return len(rng)
}

// Less check of range
func (rng Range) Less(i, j int) bool {
	return rng[i] < rng[j]
}

// Swap elements
func (rng Range) Swap(i, j int) {
	rng[i], rng[j] = rng[j], rng[i]
}

// Sort range
func (rng *Range) Sort() {
	sort.Sort(*rng)
}

// Parse of numeric ranges from string
func Parse(s string) (Range, error) {
	rng := Range{}
	err := rng.Parse(s)
	if err != nil {
		return nil, err
	}
	return rng, nil
}

package adventure

import (
	"strings"
	"testing"

	"github.com/ayzatziko/algos"
)

func TestKeys(t *testing.T) {
	in := "3 3\n1 0\n3 2\n2 1\n2 3\n4 1\n4 2\n3 0\n"
	r := strings.NewReader(in)
	keys := ReadKeys(r)

	collectCases := []struct {
		x    int
		keys []int
	}{
		{1, []int{0}},
		{0, nil},
		{4, nil},
		{2, []int{1}},
		{3, []int{2}},
	}

	for i, c := range collectCases {
		kk := CollectKeys(keys, c.x)
		if !sameSet(kk, c.keys) {
			t.Fatalf("%d: expected %v is equal to %v\n", i, kk, c.keys)
		}
	}

	checkCases := []struct {
		x    int
		keys []bool
		ok   bool
	}{
		// fail
		{4, []bool{false, true}, false},
		{4, []bool{false, false, true}, false},
		{4, []bool{true}, false},
		{3, []bool{false, true}, false},

		// success
		{0, nil, true},
		{0, []bool{true}, true},
		{4, []bool{false, true, true}, true},
		{4, []bool{true, true, true}, true},
		{3, []bool{true}, true},
	}

	for i, c := range checkCases {
		ok := CheckKeys(keys, c.x, c.keys)
		if ok != c.ok {
			if !c.ok {
				t.Fatalf("%d: expected failure, x: %d, keys: %v, g: %s\n", i, c.x, c.keys, algos.String(keys.req))
			} else {
				t.Fatalf("%d: expected success, x: %d, keys: %v, g: %s\n", i, c.x, c.keys, algos.String(keys.req))
			}
		}
	}
}

// sameSet returns true if the first slice contains all elements from the second
func sameSet(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if len(a) != len(b) {
		return false
	}

	contains := func(s []int, i int) bool {
		for _, v := range s {
			if v == i {
				return true
			}
		}

		return false
	}

	for range a {
		for _, j := range b {
			if !contains(a, j) {
				return false
			}
		}
	}

	for range b {
		for _, j := range a {
			if !contains(b, j) {
				return false
			}
		}
	}

	return true
}

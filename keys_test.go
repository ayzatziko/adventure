package adventure

import (
	"testing"
)

func TestKeys(t *testing.T) {
	keys := Keys{
		nodeKeys: map[int][]int{
			1: []int{0}, 2: []int{1}, 3: []int{2},
		},
		req: map[int][]int{
			4: []int{1, 2},
			3: []int{0},
		},
	}

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
		kk := CollectKeys(&keys, c.x)
		if !same(kk, c.keys) || !same(c.keys, kk) {
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
		ok := CheckKeys(&keys, c.x, c.keys)
		if ok != c.ok {
			if !c.ok {
				t.Fatalf("%d: expected failure, x: %d, keys: %v, g: %v\n", i, c.x, c.keys, keys.req)
			} else {
				t.Fatalf("%d: expected success, x: %d, keys: %v, g: %v\n", i, c.x, c.keys, keys.req)
			}
		}
	}
}

func same(a, b []int) bool {
	containes := func(ints []int, v int) bool {
		for _, i := range ints {
			if i == v {
				return true
			}
		}

		return false
	}

	for _, av := range a {
		if !containes(b, av) {
			return false
		}
	}

	return true
}

package adventure

import (
	"io"

	"github.com/ayzatziko/algos"
)

type Keys struct {
	pos *algos.Graph
	req *algos.Graph
}

func ReadKeys(r io.Reader) *Keys {
	posg := algos.ReadGraph(r)
	reqg := algos.ReadGraph(r)
	return &Keys{posg, reqg}
}

func CollectKeys(k *Keys, x int) []int {
	return algos.Paths(k.pos, x)
}

func CheckKeys(k *Keys, x int, keys []bool) bool {
	reqKeys := algos.Paths(k.req, x)
	if reqKeys == nil {
		return true
	}

	var gotKeys []int
	for key, ok := range keys {
		if ok {
			gotKeys = append(gotKeys, key)
		}
	}

	contains := func(s []int, i int) bool {
		for _, v := range s {
			if v == i {
				return true
			}
		}

		return false
	}

	for range gotKeys {
		for _, j := range reqKeys {
			if !contains(gotKeys, j) {
				return false
			}
		}
	}
	
	return true
}

package adventure

type Keys struct {
	nodeKeys map[int][]int
	req      map[int][]int
}

func CollectKeys(k *Keys, x int) []int {
	return k.nodeKeys[x]
}

func CheckKeys(k *Keys, x int, keys []bool) bool {
	rkeys := k.req[x]
	for _, key := range rkeys {
		if !keys[key] {
			return false
		}
	}

	return true
}

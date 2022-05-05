package p1046laststoneweightkjo

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		if stones[n-1] == stones[n-2] {
			stones = stones[:n-2]
		} else {
			stones[n-2] = stones[n-1] - stones[n-2]
			stones = stones[:n-1]
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}

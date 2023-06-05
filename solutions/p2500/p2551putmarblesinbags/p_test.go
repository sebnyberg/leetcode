package p2551putmarblesinbags

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	barCost := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		barCost[i] = weights[i] + weights[i+1]
	}
	sort.Ints(barCost)
	minCost := weights[0] + weights[n-1]
	maxCost := minCost
	for i := 0; i < k-1; i++ {
		minCost += barCost[i]
		maxCost += barCost[n-2-i]
	}
	return int64(maxCost - minCost)
}

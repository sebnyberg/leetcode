package p2262totalappealofastring

import "math"

func minimumCardPickup(cards []int) int {
	recentPos := make(map[int]int)

	getPos := func(x int) int {
		if _, exists := recentPos[x]; !exists {
			return -1
		}
		return recentPos[x]
	}

	minDist := math.MaxInt32
	for i, c := range cards {
		if j := getPos(c); j != -1 {
			minDist = min(minDist, i-j+1)
		}
		recentPos[c] = i
	}
	if minDist == math.MaxInt32 {
		return -1
	}
	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0846handofstraights

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}
	numCount := make(map[int]int)
	for _, h := range hand {
		numCount[h]++
	}
	sort.Ints(hand)
	// For each unique value with a non-zero count, create a hand
	for _, h := range hand {
		if numCount[h] == 0 {
			continue
		}
		for x := 0; x < groupSize; x++ {
			if numCount[h+x] == 0 {
				return false
			}
			numCount[h+x]--
		}
	}
	return true
}

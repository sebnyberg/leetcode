package p3074appleredistributionintoboxes

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	var appleSum int
	for _, x := range apple {
		appleSum += x
	}
	var j int
	for appleSum >= 1 {
		appleSum -= capacity[j]
		j++
	}
	return j
}

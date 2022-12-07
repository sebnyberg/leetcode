package p1051heightchecker

import "sort"

func heightChecker(heights []int) int {
	n := len(heights)
	a := make([]int, n)
	copy(a, heights)
	sort.Ints(a)
	var res int
	for i := range a {
		if a[i] != heights[i] {
			res++
		}
	}
	return res
}

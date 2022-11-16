package p0977squaresofasortedarray

import "sort"

func sortedSquares(nums []int) []int {
	var res []int
	for _, x := range nums {
		res = append(res, x*x)
	}
	sort.Ints(res)
	return res
}

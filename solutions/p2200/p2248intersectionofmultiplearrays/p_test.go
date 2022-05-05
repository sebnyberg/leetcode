package p2248intersectionofmultiplearrays

import "sort"

func intersection(nums [][]int) []int {
	numCount := make(map[int]int)
	for _, ns := range nums {
		for _, x := range ns {
			numCount[x]++
		}
	}
	var res []int
	for num, count := range numCount {
		if count == len(nums) {
			res = append(res, num)
		}
	}
	sort.Ints(res)
	return res
}

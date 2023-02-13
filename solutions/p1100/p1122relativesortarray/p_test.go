package p1122relativesortarray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for _, x := range arr1 {
		m[x]++
	}
	res := make([]int, 0, len(arr1))
	for _, x := range arr2 {
		for i := 0; i < m[x]; i++ {
			res = append(res, x)
		}
		delete(m, x)
	}
	var rest []int
	for x := range m {
		for i := 0; i < m[x]; i++ {
			rest = append(rest, x)
		}
	}
	sort.Ints(rest)
	res = append(res, rest...)
	return res
}

package p2363mergesimilaritems

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	items := make(map[int]int)
	for _, v := range items2 {
		items[v[0]] += v[1]
	}
	for _, v := range items1 {
		items[v[0]] += v[1]
	}
	res := make([][]int, 0)
	for v, w := range items {
		res = append(res, []int{v, w})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

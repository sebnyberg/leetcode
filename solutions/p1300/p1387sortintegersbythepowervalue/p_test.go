package p1387sortintegersbythepowervalue

import "sort"

func getKth(lo int, hi int, k int) int {
	sim := func(x int) int {
		var moves int
		for x > 1 {
			if x%2 == 1 {
				x = x*3 + 1
			} else {
				x /= 2
			}
			moves++
		}
		return moves
	}
	res := make([][2]int, hi-lo+1)
	for x := lo; x <= hi; x++ {
		res[x-lo] = [2]int{x, sim(x)}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][1] == res[j][1] {
			return res[i][0] < res[j][0]
		}
		return res[i][1] < res[j][1]
	})
	return res[k-1][0]
}

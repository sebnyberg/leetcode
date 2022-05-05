package p1337kweakestrows

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	rowvals := make([][2]int, len(mat))
	for i, row := range mat {
		rowvals[i][0] = i
		for _, n := range row {
			rowvals[i][1] += n & 1
		}
	}
	sort.Slice(rowvals, func(i, j int) bool {
		if rowvals[i][1] == rowvals[j][1] {
			return rowvals[i][0] < rowvals[j][0]
		}
		return rowvals[i][1] < rowvals[j][1]
	})
	res := make([]int, 0)
	for i := range rowvals[:k] {
		res = append(res, rowvals[i][0])
	}
	return res
}

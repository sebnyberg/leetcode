package p1380luckynumbersinamatrix

import "math"

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	min := make([]int, m)
	max := make([]int, n)
	for i := range min {
		min[i] = math.MaxInt32
	}
	for i := range matrix {
		for j, v := range matrix[i] {
			if v < min[i] {
				min[i] = v
			}
			if v > max[j] {
				max[j] = v
			}
		}
	}
	var res []int
	for i := range matrix {
		for j, v := range matrix[i] {
			if v == min[i] && v == max[j] {
				res = append(res, v)
			}
		}
	}
	return res
}

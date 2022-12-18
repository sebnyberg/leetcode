package p2500deletegreatestvalueineachrow

import "sort"

func deleteGreatestValue(grid [][]int) int {
	var sum int
	for i := range grid {
		sort.Ints(grid[i])
	}
	for j := range grid[0] {
		var maxVal int
		for i := range grid {
			maxVal = max(maxVal, grid[i][j])
		}
		sum += maxVal
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

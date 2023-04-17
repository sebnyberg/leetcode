package p2639findthewidthofcolumnsofagrid

import "fmt"

func findColumnWidth(grid [][]int) []int {
	n := len(grid[0])
	res := make([]int, n)
	for i := range grid {
		for j := range grid[i] {
			res[j] = max(res[j], len(fmt.Sprint(grid[i][j])))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

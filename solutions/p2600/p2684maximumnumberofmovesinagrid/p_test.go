package p2684maximumnumberofmovesinagrid

import "math"

func maxMoves(grid [][]int) int {
	m := len(grid)
	prev := make([]int, m)
	curr := make([]int, m)
	var res int
	for i := range len(grid) {
		prev[i] = 0 // all values in the first column have a length of 0
	}
	for j := 1; j < len(grid[0]); j++ {
		for i := range curr {
			curr[i] = math.MinInt32
		}
		for i := range grid {
			if i > 0 && grid[i-1][j-1] < grid[i][j] {
				curr[i] = prev[i-1] + 1
			}
			if grid[i][j-1] < grid[i][j] {
				curr[i] = max(curr[i], prev[i]+1)
			}
			if i < len(grid)-1 && grid[i+1][j-1] < grid[i][j] {
				curr[i] = max(curr[i], prev[i+1]+1)
			}
			res = max(res, curr[i])
		}
		prev, curr = curr, prev
	}
	return res
}

package p1289minimumfallingpathsumii

import "math"

func minFallingPathSum(grid [][]int) int {
	// The "difficult" part is to calculate the minimum value while excluding a
	// certain value. However, the grid is very small (<=200) so a O(n^2) isn't
	// too bad.
	n := len(grid)
	vals := make([]int, n)
	zero := make([]int, n)
	for i := range zero {
		zero[i] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		copy(vals, zero)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				vals[k] = min(vals[k], grid[i-1][j])
			}
		}
		for j := range grid[i] {
			grid[i][j] += vals[j]
		}
	}
	res := math.MaxInt32
	for _, x := range grid[n-1] {
		res = min(res, x)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

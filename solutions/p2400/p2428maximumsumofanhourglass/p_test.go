package p2428maximumsumofanhourglass

func maxSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var res int
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			sum := grid[i][j] + grid[i][j+1] + grid[i][j+2] +
				grid[i+1][j+1] +
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			res = max(res, sum)
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

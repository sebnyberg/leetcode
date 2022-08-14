package p2373largestlocalvaluesinamatrix

func largestLocal(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m-2)
	for i := range res {
		res[i] = make([]int, n-2)
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			var maxVal int
			for k := i - 1; k <= i+1; k++ {
				for kk := j - 1; kk <= j+1; kk++ {
					if x := grid[k][kk]; x > maxVal {
						maxVal = x
					}
				}
			}
			res[i-1][j-1] = maxVal
		}
	}
	return res
}

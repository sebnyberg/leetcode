package p0931minimumfallingpathsum

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			d := matrix[i-1][j]
			if j > 0 {
				d = min(d, matrix[i-1][j-1])
			}
			if j < len(matrix[i])-1 {
				d = min(d, matrix[i-1][j+1])
			}
			matrix[i][j] += d
		}
	}
	res := matrix[m-1][0]
	for _, v := range matrix[m-1] {
		res = min(res, v)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

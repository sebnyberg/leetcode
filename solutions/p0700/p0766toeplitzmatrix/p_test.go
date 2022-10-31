package p0766toeplitzmatrix

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}

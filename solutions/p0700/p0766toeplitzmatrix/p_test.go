package p0766toeplitzmatrix

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for col := 0; col < n; col++ {
		for k := 1; col+k < n && k < m; k++ {
			if matrix[k][col+k] != matrix[k-1][col+k-1] {
				return false
			}
		}
	}
	for row := 1; row < m; row++ {
		for k := 1; row+k < m && k < n; k++ {
			if matrix[row+k][k] != matrix[row+k-1][k-1] {
				return false
			}
		}
	}
	return true
}

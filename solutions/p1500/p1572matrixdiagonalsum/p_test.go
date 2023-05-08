package p1572matrixdiagonalsum

func diagonalSum(mat [][]int) int {
	var sum int
	n := len(mat)
	for i := range mat {
		sum += mat[i][i] + mat[i][n-i-1]
	}
	if n&1 == 1 {
		sum -= mat[n/2][n/2]
	}
	return sum
}

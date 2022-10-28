package p0867transposematrix

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
		for j := range matrix {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}

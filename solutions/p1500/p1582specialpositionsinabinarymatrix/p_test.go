package p1582specialpositionsinabinarymatrix

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i := range mat {
		for j := range mat[i] {
			rows[i] += mat[i][j]
			cols[j] += mat[i][j]
		}
	}
	var res int
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				res++
			}
		}
	}
	return res
}

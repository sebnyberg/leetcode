package p1252cellswithoddvaluesinamatrix

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, x := range indices {
		rows[x[0]]++
		cols[x[1]]++
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += (rows[i] + cols[j]) & 1
		}
	}
	return res
}

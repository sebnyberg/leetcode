package p2713maximumstrictlyincreasingcellsinamatrix

import "sort"

func maxIncreasingCells(mat [][]int) int {
	// There are A LOT of edges in the matrix.
	// Working backwards is as good/bad as working forwards.
	// What if we store the maximum possible path length per row and column
	// given a certain iteration. Then for each unique value, the value of all
	// rows/columns that the values are in will be updated to be
	// max(row+1,col+1).
	//
	// For each value, collect its coordinate
	pos := make(map[int][][2]int)
	for i := range mat {
		for j := range mat[i] {
			pos[mat[i][j]] = append(pos[mat[i][j]], [2]int{i, j})
		}
	}
	// Collect unique values
	vals := []int{}
	for v := range pos {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)

	next := [][]int{}
	for _, v := range vals {
		next = next[:0]
		for _, x := range pos[v] {
			// push an update of {row, col, value}
			next = append(next,
				[]int{x[0], x[1], 1 + max(rows[x[0]], cols[x[1]])},
			)
		}
		for _, upd := range next {
			row := upd[0]
			col := upd[1]
			val := upd[2]
			rows[row] = max(rows[row], val)
			cols[col] = max(cols[col], val)
		}
	}
	var res int
	for _, v := range rows {
		res = max(res, v)
	}
	for _, v := range cols {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

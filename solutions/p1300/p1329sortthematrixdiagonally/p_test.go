package p1329sortthematrixdiagonally

import "sort"

type diagonal struct {
	mat  [][]int
	i, j int
}

func (d *diagonal) Less(i, j int) bool {
	return d.mat[d.i+i][d.j+i] < d.mat[d.i+j][d.j+j]
}

func (d *diagonal) Swap(i, j int) {
	i1, j1 := d.i+i, d.j+i
	i2, j2 := d.i+j, d.j+j
	d.mat[i1][j1], d.mat[i2][j2] = d.mat[i2][j2], d.mat[i1][j1]
}

func (d *diagonal) Len() int {
	m := len(d.mat)
	n := len(d.mat[0])
	return min(n-d.j, m-d.i)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diagonalSort(mat [][]int) [][]int {
	for i := range mat {
		sort.Sort(&diagonal{mat, i, 0})
	}
	n := len(mat[0])
	for j := 1; j < n; j++ {
		sort.Sort(&diagonal{mat, 0, j})
	}
	return mat
}

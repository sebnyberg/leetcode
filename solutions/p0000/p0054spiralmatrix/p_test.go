package p0054spiralmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_spiralOrder(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, spiralOrder(tc.matrix))
		})
	}
}

type direction int

func (d direction) di() int {
	switch d % 4 {
	case 1:
		return 1
	case 3:
		return -1
	}
	return 0
}

func (d direction) dj() int {
	switch d % 4 {
	case 0:
		return 1
	case 2:
		return -1
	}
	return 0
}

func spiralOrder(matrix [][]int) []int {
	nilEntry := -101
	m, n := len(matrix), len(matrix[0])
	valid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && matrix[i][j] != nilEntry
	}
	var i, j int
	var dir direction
	res := make([]int, 0, m*n)
	for valid(i, j) {
		res = append(res, matrix[i][j])
		matrix[i][j] = nilEntry
		if !valid(i+dir.di(), j+dir.dj()) {
			dir++
		}
		i += dir.di()
		j += dir.dj()
	}
	return res
}

// func spiralOrder(matrix [][]int) []int {
// 	// -100 is the smallest number for an entry in the matrix
// 	nilEntry := -101
// 	// Pad the matrix with nil entries on the left / right / below to make
// 	// traversal less complex in code
// 	// This is not needed for the first row since it will be filled
// 	// with nilEntries anyway
// 	m := len(matrix[0])
// 	n := len(matrix)
// 	res := make([]int, 0, m*n)
// 	m += 2
// 	n++

// 	for i := range matrix {
// 		matrix[i] = append(matrix[i], nilEntry, 0)
// 		copy(matrix[i][1:], matrix[i])
// 		matrix[i][0] = nilEntry
// 	}
// 	matrix = append(matrix, make([]int, m))
// 	for i := range matrix[n-1] {
// 		matrix[n-1][i] = nilEntry
// 	}

// 	// Iterate through the matrix
// 	i, j := 0, 1
// 	for {
// 		// right
// 		if matrix[i][j] == nilEntry {
// 			return res
// 		}
// 		for matrix[i][j] != nilEntry {
// 			res = append(res, matrix[i][j])
// 			matrix[i][j] = nilEntry
// 			j++
// 		}
// 		j--
// 		i++
// 		// down
// 		if matrix[i][j] == nilEntry {
// 			return res
// 		}
// 		for matrix[i][j] != nilEntry {
// 			res = append(res, matrix[i][j])
// 			matrix[i][j] = nilEntry
// 			i++
// 		}
// 		i--
// 		j--
// 		// left
// 		if matrix[i][j] == nilEntry {
// 			return res
// 		}
// 		for matrix[i][j] != nilEntry {
// 			res = append(res, matrix[i][j])
// 			matrix[i][j] = nilEntry
// 			j--
// 		}
// 		j++
// 		i--
// 		// up
// 		if matrix[i][j] == nilEntry {
// 			return res
// 		}
// 		for matrix[i][j] != nilEntry {
// 			res = append(res, matrix[i][j])
// 			matrix[i][j] = nilEntry
// 			i--
// 		}
// 		i++
// 		j++
// 	}
// }

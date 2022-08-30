package p0048rotateimage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotate(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			rotate(tc.matrix)
			require.Equal(t, tc.want, tc.matrix)
		})
	}
}

func rotate(matrix [][]int) {
	// There are two options:
	//
	// 1. Flip along y axis, then flip along anti-diagonal
	// 2. Flip along x axis, then flip along diagonal
	//
	// Flipping along the diagonal is bad for the CPU cache no matter which
	// direction. However, when it comes to flipping the y versus x axis, the
	// y-axis is more cache friendly (data is laid out per-row in memory), so
	// let's go with option 1.
	//
	for i := range matrix {
		for l, r := 0, len(matrix)-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
	// Flip along anti-diagonal
	n := len(matrix[0])
	for i := range matrix {
		k := n - i - 1
		for j := 0; j < k; j++ {
			col := n - i - 1
			row := n - j - 1
			matrix[i][j], matrix[row][col] = matrix[row][col], matrix[i][j]
		}
	}
}

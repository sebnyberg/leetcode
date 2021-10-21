package p0566reshapethematrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matrixReshape(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, matrixReshape(tc.mat, tc.r, tc.c))
		})
	}
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// The user wants 'r' rows and 'c' cols in the output
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	out := make([][]int, r)
	for i := range out {
		out[i] = make([]int, c)
	}
	var pos int
	for i := range mat {
		for j := range mat[i] {
			out[pos/c][pos%c] = mat[i][j]
			pos++
		}
	}
	return out
}

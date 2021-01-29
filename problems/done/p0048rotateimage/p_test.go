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
	// To achieve a rotation, transpose on the diagonal, then swap along an axis
	// Transposition on the \ diagonal results in a swap along the
	// x-axis which is more cache-efficient
	n := len(matrix)
	for r := 0; r < n-1; r++ {
		for c := r + 1; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			matrix[r][c], matrix[r][n-c-1] = matrix[r][n-c-1], matrix[r][c]
		}
	}
}

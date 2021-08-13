package p0073matrixzeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setZeroes(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{[][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}, [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			setZeroes(tc.matrix)
			require.Equal(t, tc.want, tc.matrix)
		})
	}
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	cols := make([]bool, n)
	rows := make([]bool, m)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

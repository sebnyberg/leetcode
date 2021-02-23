package search2dmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchMatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.matrix, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, searchMatrix(tc.matrix, tc.target))
		})
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	// Naive search
	for i := 0; i < len(matrix) && matrix[i][0] <= target; i++ {
		for j := 0; j < len(matrix[i]) && matrix[i][j] <= target; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

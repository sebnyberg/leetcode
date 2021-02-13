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
	m := len(matrix)
	n := len(matrix[0])

	// Check first col
	var shouldZeroFirstCol bool
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			shouldZeroFirstCol = true
			break
		}
	}

	// Check first row
	// Whether first row should be zeroed is stored in matrix[0][0]
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			matrix[0][0] = 0
			break
		}
	}

	// If a value is zero, set the corresponding first value
	// in the same col / row to zero. This is a non-destructive
	// maneuver since we have already stored whether the first
	// col / row should all be zeroed out in the "fistnum" value
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Zero out entries in rows for which the first value is zero
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}

	// Zero out entries in columns for which the first value is zero
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := range matrix {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if shouldZeroFirstCol {
		for j := 0; j < m; j++ {
			matrix[j][0] = 0
		}
	}
}

package p1074nsubmatricesthatsumtotarget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		want   int
	}{
		{[][]int{{1, -1}, {-1, 1}}, 0, 5},
		{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
		{[][]int{{904}}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, numSubmatrixSumTarget(tc.matrix, tc.target))
		})
	}
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// First, count the sums of each row block in the matrix
	m, n := len(matrix), len(matrix[0])
	prefixSum := make([][]int, m+1)
	prefixSum[0] = make([]int, n+1)
	for row := 1; row <= m; row++ {
		prefixSum[row] = make([]int, n+1)
		for col := 1; col <= n; col++ {
			prefixSum[row][col] += prefixSum[row-1][col] + prefixSum[row][col-1]
			prefixSum[row][col] -= prefixSum[row-1][col-1]
			prefixSum[row][col] += matrix[row-1][col-1]
		}
	}

	res := 0
	for startRow := 1; startRow <= m; startRow++ {
		for endRow := startRow; endRow <= m; endRow++ {
			for startCol := 1; startCol <= n; startCol++ {
				for endCol := startCol; endCol <= n; endCol++ {
					val := prefixSum[endRow][endCol]
					val -= prefixSum[endRow][startCol-1]
					val -= prefixSum[startRow-1][endCol]
					val += prefixSum[startRow-1][startCol-1]
					if val == target {
						res++
					}
				}
			}
		}
	}

	return res
}

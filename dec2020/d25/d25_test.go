package d25_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDiagonalOrder(t *testing.T) {
	for _, tc := range []struct {
		in   [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
	} {
		t.Run(fmt.Sprintf("+%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, findDiagonalOrder(tc.in))
		})
	}
}

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	res := make([]int, 0, m*n)
	var i, j int
	// Seed with first element (0,0)
	res = append(res, matrix[i][j])
	// Iterate over diagonals (N+M-1 total)
	for diag := 1; diag < n+m-1; diag++ {
		// Uneven diagonals, down/left
		if diag%2 == 1 {
			// Move right if possible, else move down
			if j < n-1 {
				j++
			} else {
				i++
			}
			// Move down/left until reaching the edge
			for {
				res = append(res, matrix[i][j])
				if i == m-1 || j == 0 {
					break
				}
				i++
				j--
			}

			continue
		}
		// Uneven diagonal, up/right
		// Move down if possible, else move right
		if i < m-1 {
			i++
		} else {
			j++
		}

		// Move up/right until reaching the edge
		for {
			res = append(res, matrix[i][j])
			if i == 0 || j == n-1 {
				break
			}
			i--
			j++
		}
	}
	return res
}

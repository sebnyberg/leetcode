package p1277countsquaresubmatriceswithallones

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countSquares(t *testing.T) {
	for i, tc := range []struct {
		matrix [][]int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[0,1,1,1],[1,1,1,1],[0,1,1,1]]"),
			15,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countSquares(tc.matrix))
		})
	}
}

func countSquares(matrix [][]int) int {
	// We can use a 2d- prefix sum and just count for each possible position.
	m := len(matrix)
	n := len(matrix[0])

	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i := range matrix {
		for j, v := range matrix[i] {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}

	var res int
	for k := 1; k <= len(matrix); k++ {
		want := k * k
		for i := k; i < m+1; i++ {
			for j := k; j < n+1; j++ {
				got := pre[i][j] - pre[i-k][j] - pre[i][j-k] + pre[i-k][j-k]
				if want == got {
					res++
				}
			}
		}
	}
	return res
}

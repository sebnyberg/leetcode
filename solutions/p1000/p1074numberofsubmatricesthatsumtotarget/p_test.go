package p1074numberofsubmatricesthatsumtotarget

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	for i, tc := range []struct {
		matrix [][]int
		target int
		want   int
	}{
		{leetcode.ParseMatrix("[[0,1,0],[1,1,1],[0,1,0]]"), 0, 4},
		{leetcode.ParseMatrix("[[1,-1],[-1,1]]"), 0, 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numSubmatrixSumTarget(tc.matrix, tc.target))
		})
	}
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = matrix[i-1][j-1] + pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1]
		}
	}
	var res int
	for i := range matrix {
		for j := range matrix[i] {
			for ii := 0; ii <= i; ii++ {
				for jj := 0; jj <= j; jj++ {
					sum := pre[i+1][j+1] + pre[ii][jj]
					sum -= pre[i+1][jj]
					sum -= pre[ii][j+1]
					if sum == target {
						res++
					}
				}
			}
		}
	}
	return res
}

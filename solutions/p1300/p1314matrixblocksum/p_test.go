package p1314matrixblocksum

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_matrixBlockSum(t *testing.T) {
	for i, tc := range []struct {
		mat  [][]int
		k    int
		want [][]int
	}{
		{
			leetcode.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"),
			1,
			leetcode.ParseMatrix("[[12,21,16],[27,45,33],[24,39,28]]"),
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, matrixBlockSum(tc.mat, tc.k))
		})
	}
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	// 2d-prefix sum
	m := len(mat)
	n := len(mat[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}
	for i := range mat {
		for j, v := range mat[i] {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}
	squareVal := func(i1, j1, i2, j2 int) int {
		i2++
		j2++
		return pre[i2][j2] - pre[i1][j2] - pre[i2][j1] + pre[i1][j1]
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			i1 := max(0, i-k)
			j1 := max(0, j-k)
			i2 := min(i+k, m-1)
			j2 := min(j+k, n-1)
			res[i][j] = squareVal(i1, j1, i2, j2)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

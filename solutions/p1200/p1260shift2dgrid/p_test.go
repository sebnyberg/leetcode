package p1260shift2dgrid

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shiftGrid(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		k    int
		want [][]int
	}{
		{
			leetcode.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"),
			1,
			leetcode.ParseMatrix("[[9,1,2],[3,4,5],[6,7,8]]"),
		},
		{
			leetcode.ParseMatrix("[[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]"),
			4,
			leetcode.ParseMatrix("[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]"),
		},
		{
			leetcode.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"),
			9,
			leetcode.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"),
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, shiftGrid(tc.grid, tc.k))
		})
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k = m*n - k%(m*n)
	res := make([][]int, m)
	for i := range grid {
		res[i] = make([]int, n)
		for j := range grid[0] {
			res[i][j] = grid[(k/n)%m][k%n]
			k++
		}
	}
	return res
}

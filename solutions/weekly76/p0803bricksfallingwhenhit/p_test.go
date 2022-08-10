package p4

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_hitBricks(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		hits [][]int
		want []int
	}{
		{
			leetcode.ParseMatrix("[[1,0,0,0],[1,1,1,0]]"),
			leetcode.ParseMatrix("[[1,0]]"),
			[]int{2},
		},
		{
			leetcode.ParseMatrix("[[1,0,0,0],[1,1,0,0]]"),
			leetcode.ParseMatrix("[[1,1],[1,0]]"),
			[]int{0, 0},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, hitBricks(tc.grid, tc.hits))
		})
	}
}

func hitBricks(grid [][]int, hits [][]int) []int {
}

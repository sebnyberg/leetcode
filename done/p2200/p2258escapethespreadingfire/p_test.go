package p2

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countUnguarded(t *testing.T) {
	for _, tc := range []struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
		want   int
	}{
		{
			4, 6,
			leetcode.ParseMatrix("[[0,0],[1,1],[2,3]]"),
			leetcode.ParseMatrix("[[0,1],[2,2],[1,4]]"),
			7,
		},
		{
			3, 3,
			leetcode.ParseMatrix("[[1,1]]"),
			leetcode.ParseMatrix("[[0,1],[1,0],[2,1],[1,2]]"),
			7,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, countUnguarded(tc.m, tc.n, tc.guards, tc.walls))
		})
	}
}

type fovCell struct {
	i, j   int
	dirIdx int
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	// Trying each guard and each direction is too expensive.
	// Having a grid of all positions is OK
	// When encountering a certain direction / cell twice, stop traversal
	seen := make(map[fovCell]struct{}, m*n)
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	blocked := make(map[[2]int]struct{}, len(walls))
	guarded := make(map[[2]int]struct{}, m*n)
	for _, w := range walls {
		blocked[[2]int{w[0], w[1]}] = struct{}{}
	}
	for _, g := range guards {
		guarded[[2]int{g[0], g[1]}] = struct{}{}
	}
	ok := func(i, j, dirIdx int) bool {
		_, exists := seen[fovCell{i, j, dirIdx}]
		_, blocked := blocked[[2]int{i, j}]
		return i >= 0 && j >= 0 && i < m && j < n && !blocked && !exists
	}
	for _, g := range guards {
		for dirIdx, d := range dirs {
			i, j := g[0], g[1]
			for {
				i += d[0]
				j += d[1]
				// up/down is the same thing
				if !ok(i, j, dirIdx&1) {
					break
				}
				seen[fovCell{i, j, dirIdx & 1}] = struct{}{}
				guarded[[2]int{i, j}] = struct{}{}
			}
		}
	}
	res := m*n - len(guarded) - len(blocked)
	return res
}

package p1091shortpathinbinmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{0, 1}, {1, 0}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			got := shortestPathBinaryMatrix(tc.grid)
			require.Equal(t, tc.want, got)
		})
	}
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid[0])
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	curr := [][2]int{{0, 0}}
	next := [][2]int{}
	seen := make(map[[2]int]bool)
	seen[curr[0]] = true
	dirs := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < n && j < n
	}

	for k := 2; len(curr) > 0; k++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				q := [2]int{ii, jj}
				if !ok(ii, jj) || seen[q] || grid[ii][jj] == 1 {
					continue
				}
				if q == [2]int{n - 1, n - 1} {
					return k
				}
				seen[q] = true
				next = append(next, q)
			}
		}
		curr, next = next, curr
	}

	return -1
}

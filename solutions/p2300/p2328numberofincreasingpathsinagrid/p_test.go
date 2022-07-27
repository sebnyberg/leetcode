package p2328numberofincreasingpathsinagrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPaths(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 1}, {3, 4}}, 8},
		{[][]int{{1}, {2}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, countPaths(tc.grid))
		})
	}
}

func countPaths(grid [][]int) int {
	// For each position, perform dp with memoization
	m := len(grid)
	n := len(grid[0])
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var res int
	for i := range grid {
		for j := range grid[i] {
			res = (res + dp(mem, grid, i, j, m, n)) % mod
		}
	}
	return res % mod
}

const mod = 1e9 + 7

var dirs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func dp(mem, grid [][]int, i, j, m, n int) int {
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	ok := func(ii, jj int) bool {
		return ii >= 0 && jj >= 0 && ii < m && jj < n && grid[ii][jj] < grid[i][j]
	}

	res := 1
	for _, d := range dirs {
		ii, jj := i+d[0], j+d[1]
		if ok(ii, jj) {
			res = (res + dp(mem, grid, ii, jj, m, n)) % mod
		}
	}

	mem[i][j] = res
	return res
}

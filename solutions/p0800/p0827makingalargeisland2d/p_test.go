package p4

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_largestIsland(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[0,0],[0,0]]"), 1},
		{leetcode.ParseMatrix("[[1,0],[0,1]]"), 3},
		{leetcode.ParseMatrix("[[1,1],[1,0]]"), 4},
		{leetcode.ParseMatrix("[[1,1],[1,1]]"), 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, largestIsland(tc.grid))
		})
	}
}

func largestIsland(grid [][]int) int {
	// typical DSU exercise
	m := len(grid)
	n := len(grid[0])
	parent := make([][][2]int, m)
	size := make([][]int, m)
	for i := range parent {
		parent[i] = make([][2]int, n)
		size[i] = make([]int, n)
		for j := range parent[i] {
			size[i][j] = 1
			parent[i][j] = [2]int{i, j}
		}
	}

	var find func(a [2]int) [2]int
	find = func(a [2]int) [2]int {
		if parent[a[0]][a[1]] == a {
			return a
		}
		root := find(parent[a[0]][a[1]])
		parent[a[0]][a[1]] = root
		return root
	}

	union := func(a, b [2]int) {
		ra := find(a)
		rb := find(b)
		if ra != rb {
			parent[ra[0]][ra[1]] = rb
			size[rb[0]][rb[1]] += size[ra[0]][ra[1]]
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// First, join all islands together
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			for _, d := range dirs {
				ii, jj := i+d[0], j+d[1]
				if !ok(ii, jj) || grid[ii][jj] == 0 {
					continue
				}
				union([2]int{i, j}, [2]int{ii, jj})
			}
		}
	}

	// Then find the best flip (if any)
	var res int
	nei := make(map[[2]int]struct{})
	for i := range grid {
		for j := range grid[i] {
			res = max(res, size[i][j])
			if grid[i][j] == 1 {
				continue
			}
			for k := range nei {
				delete(nei, k)
			}
			// Changing this to a 1 will join all islands in all four directions
			for _, d := range dirs {
				ii, jj := i+d[0], j+d[1]
				if !ok(ii, jj) || grid[ii][jj] == 0 {
					continue
				}
				root := find([2]int{ii, jj})
				if _, exists := nei[root]; exists {
					continue
				}
				nei[root] = struct{}{}
			}
			// We are joining all unique neighbouring islands together
			subRes := 1
			for r := range nei {
				subRes += size[r[0]][r[1]]
			}
			res = max(res, subRes)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

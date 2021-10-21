package p0827makingalargeisland

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestIsland(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0}, {0, 1}}, 3},
		{[][]int{{1, 1}, {1, 0}}, 4},
		{[][]int{{1, 1}, {1, 1}}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, largestIsland(tc.grid))
		})
	}
}

func largestIsland(grid [][]int) int {
	// * Place 1s in groups based on their neighbours using union find
	// * For each zero in the grid, calculate the size of all unique groups
	//   reachable from that zero. The largest such grouping of unique islands
	//   is the answer.
	n := len(grid)
	parent := make([]int, n*n)
	size := make([]int, n*n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(a int) int {
		root := parent[a]
		for root != parent[root] {
			root = parent[root]
		}
		return root
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
			size[rb] += size[ra]
		}
	}
	pos := func(i, j int) int {
		return n*i + j
	}

	// Add existing islands to DSU
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			for _, near := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
			} {
				ni, nj := near[0], near[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= n || grid[ni][nj] != 1 {
					continue
				}
				union(pos(i, j), pos(ni, nj))
			}
		}
	}

	var maxSize int
	for _, sz := range size {
		maxSize = max(maxSize, sz)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				continue
			}
			joinedSize := 1
			seen := make(map[int]struct{}, 4)
			for _, near := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
			} {
				ni, nj := near[0], near[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= n || grid[ni][nj] != 1 {
					continue
				}
				root := find(pos(ni, nj))
				if _, exists := seen[root]; !exists {
					joinedSize += size[root]
					seen[root] = struct{}{}
				}
			}
			maxSize = max(maxSize, joinedSize)
		}
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

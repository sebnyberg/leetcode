package p1368minimumcosttomakeatleastonevalidpathinagrid

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[1,1,3],[3,2,2],[1,1,4]]"),
			0,
		},
		{
			leetcode.ParseMatrix("[[1,2],[4,3]]"),
			1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.grid))
		})
	}
}

func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && !seen[i][j]
	}
	curr := [][2]int{{0, 0}}
	next := [][2]int{}
	for steps := 0; ; steps++ {
		for i := 0; i < len(curr); i++ {
			// Use BFS to find all reachable squares
			x := curr[i]
			if x == [2]int{m - 1, n - 1} {
				return steps
			}
			i := x[0]
			j := x[1]
			switch grid[i][j] {
			case 1: // right
				j++
			case 2: // left
				j--
			case 3: // down
				i++
			case 4: // up
				i--
			}
			if !ok(i, j) || seen[i][j] {
				continue
			}
			seen[i][j] = true
			curr = append(curr, [2]int{i, j})
		}
		// curr now contains all reachable nodes
		// for each reachable node, we can visit all neighbouring nodes with
		// cost 1
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				next = append(next, [2]int{ii, jj})
			}
		}
		curr, next = next, curr
	}
}

package p0490themaze

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"

	"github.com/stretchr/testify/require"
)

func Test_hasPath(t *testing.T) {
	for _, tc := range []struct {
		maze        [][]int
		start       []int
		destination []int
		want        bool
	}{
		{
			leetcode.ParseMatrix("[[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]"),
			[]int{0, 4},
			[]int{1, 2},
			true,
		},
		{
			leetcode.ParseMatrix("[[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]"),
			[]int{0, 4},
			[]int{4, 4},
			true,
		},
		{
			leetcode.ParseMatrix("[[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]"),
			[]int{0, 4},
			[]int{3, 2},
			false,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.maze), func(t *testing.T) {
			require.Equal(t, tc.want, hasPath(tc.maze, tc.start, tc.destination))
		})
	}
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	// Perform simple BFS. For each location, kick the ball in all four
	// directions, keep going until not OK. If that location is new, add it to
	// the next iteration of BFS. Rinse repeat until the destination is found,
	// or there are no more places to visit.

	m, n := len(maze), len(maze[0])
	var seen [101][101]bool
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && maze[i][j] != 1
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	seen[start[0]][start[1]] = true
	curr := [][]int{start}
	next := [][]int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, pos := range curr {
			for _, d := range dirs {
				ii, jj := pos[0], pos[1]
				for ok(ii+d[0], jj+d[1]) {
					ii += d[0]
					jj += d[1]
				}
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				if destination[0] == ii && destination[1] == jj {
					return true
				}
				seen[ii][jj] = true
				next = append(next, []int{ii, jj})
			}
		}

		curr, next = next, curr
	}

	return false
}

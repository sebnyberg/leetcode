package p1568minimumnumberofdaystodisconnectisland

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minDays(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[1,1]]"), 2},
		{leetcode.ParseMatrix("[[0,1,1,0],[0,1,1,0],[0,0,0,0]]"), 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minDays(tc.grid))
		})
	}
}

func minDays(grid [][]int) int {
	// The grid can be represented as 30 integers
	// Then just BFS and check whether there is exactly one island.
	var init [30]int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				init[i] |= 1 << j
			}
		}
	}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

	oneIsland := func(state [30]int) bool {
		// On the first occurrence of a 1 - BFS and mark 1s in a second state
		// until no more places can be found. Unset all those bits and chech
		// whether the resulting state is empty.
		var start [2]int
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if state[i]&(1<<j) > 0 {
					start = [2]int{i, j}
					goto endSearch
				}
			}
		}
		return false // no island at all

	endSearch:

		// Remove the first island that is found
		state[start[0]] &^= 1 << start[1]
		curr := [][2]int{start}
		next := [][2]int{}
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || state[ii]&(1<<jj) == 0 {
						continue
					}
					state[ii] &^= 1 << jj // unset bit
					next = append(next, [2]int{ii, jj})
				}
			}
			curr, next = next, curr
		}

		// If there was only one island, then state should be all zeroes
		for i := range state {
			if state[i] != 0 {
				return false
			}
		}

		return true
	}

	if !oneIsland(init) {
		return 0
	}

	seen := make(map[[30]int]bool)
	seen[init] = true
	curr := [][30]int{init}
	next := [][30]int{}
	for k := 1; ; k++ {
		next = next[:0]
		for _, x := range curr {
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if x[i]&(1<<j) == 0 {
						continue
					}
					// try to unset this bit and see what happens
					nextState := x
					nextState[i] &^= 1 << j
					if seen[nextState] {
						continue
					}
					seen[nextState] = true
					if !oneIsland(nextState) {
						return k
					}
					next = append(next, nextState)
				}
			}
		}
		curr, next = next, curr
	}
}

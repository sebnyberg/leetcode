package p2556disconnectpathinabinarymatrixbyatmostoneflip

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_isPossibleToCutPath(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want bool
	}{
		{
			leetcode.ParseMatrix("[[1,1,1],[1,0,0],[1,1,1]]"),
			true,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isPossibleToCutPath(tc.grid))
		})
	}
}

func isPossibleToCutPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	// First, we tag each zero-cell with top, right, bot, left depending on
	// which side it is connected to.
	state := make([][]int, m)
	for i := range state {
		state[i] = make([]int, n)
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	var curr [][2]int
	var next [][2]int
	// BFS marks connected zeroes with the provided bitmask
	bfs := func(i, j, bm int) {
		if grid[i][j] == 1 || state[i][j]&bm > 0 {
			return
		}
		curr = append(curr, [2]int{i, j})
		state[i][j] |= bm
		dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || grid[ii][jj] != 0 || state[ii][jj]&bm > 0 {
						continue
					}
					state[ii][jj] |= bm
					next = append(next, [2]int{ii, jj})
				}
			}

			curr, next = next, curr
		}
	}
	const (
		leftMark  = 1 << 0
		botMark   = 1 << 1
		rightMark = 1 << 2
		topMark   = 1 << 3
	)
	// mark each side
	for i := range grid {
		// left
		bfs(i, 0, leftMark)
		// right
		bfs(i, n-1, rightMark)
	}
	for j := range grid[0] {
		// top
		bfs(0, j, topMark)
		// bot
		bfs(m-1, j, botMark)
	}
	done := func(v int) bool {
		return (v&(leftMark|botMark)) > 0 &&
			(v&(rightMark|topMark)) > 0
	}

	// Find a cell with neighbouring cells that connect right/top to
	// bot/left
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j == 0 ||
				i == m-1 && j == n-1 {
				continue
			}
			if grid[i][j] == 0 && done(state[i][j]) {
				return true
			}
			var cellbm int
			if i == 0 {
				cellbm |= topMark
			}
			if i == m-1 {
				cellbm |= botMark
			}
			if j == 0 {
				cellbm |= leftMark
			}
			if j == n-1 {
				cellbm |= rightMark
			}
			// Try to find pair of neighbours which are both zeros and that
			// connects the right sides
			dirs := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}

			for k := range dirs {
				ii := i + dirs[k][0]
				jj := j + dirs[k][1]
				if !ok(ii, jj) || grid[ii][jj] != 0 {
					continue
				}
				cellbm |= state[ii][jj]
			}
			if done(cellbm) {
				return true
			}
		}
	}
	return false
}

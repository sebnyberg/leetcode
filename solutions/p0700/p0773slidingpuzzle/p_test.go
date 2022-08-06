package p0773slidingpuzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_slidingPuzzle(t *testing.T) {
	for _, tc := range []struct {
		boards [][]int
		want   int
	}{
		{[][]int{{1, 2, 3}, {4, 0, 5}}, 1},
		{[][]int{{1, 2, 3}, {5, 4, 0}}, -1},
		{[][]int{{4, 1, 2}, {5, 0, 3}}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boards), func(t *testing.T) {
			require.Equal(t, tc.want, slidingPuzzle(tc.boards))
		})
	}
}

// https://www.github.com/sebnyberg/leetcode
func slidingPuzzle(board [][]int) int {
	type state struct {
		board [2][3]uint8
		i, j  int8
	}

	// Parse initial state
	var init state
	for i := range board {
		for j, v := range board[i] {
			init.board[i][j] = uint8(v)
			if v == 0 {
				init.i = int8(i)
				init.j = int8(j)
			}
		}
	}

	directions := [][2]int8{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	done := [2][3]uint8{{1, 2, 3}, {4, 5, 0}}
	curr := []state{init}
	next := []state{}
	seen := make(map[state]struct{})
	seen[curr[0]] = struct{}{}

	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			if x.board == done {
				return steps
			}
			for _, d := range directions {
				ii, jj := x.i+d[0], x.j+d[1]
				if ii < 0 || jj < 0 || ii >= 2 || jj >= 3 {
					continue
				}
				// Copy board and swap tiles
				s := x
				s.board[ii][jj], s.board[x.i][x.j] = s.board[x.i][x.j], s.board[ii][jj]
				s.i = ii
				s.j = jj
				if _, exists := seen[s]; exists {
					continue
				}
				seen[s] = struct{}{}
				next = append(next, s)
			}
		}
		curr, next = next, curr
	}
	return -1
}

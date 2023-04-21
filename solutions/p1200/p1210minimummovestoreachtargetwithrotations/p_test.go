package p1210minimummovestoreachtargetwithrotations

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumMoves(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,0,0,0,0,1],[1,1,0,0,1,0],[0,0,0,0,1,1],[0,0,1,0,1,0],[0,1,1,0,0,0],[0,1,1,0,0,0]]"),
			11,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumMoves(tc.grid))
		})
	}
}

const (
	dirRight = 0
	dirDown  = 1
)

type state struct {
	pos [2][2]int
	dir uint8
}

func minimumMoves(grid [][]int) int {
	// There are not that many location/directions that the snake can be in.
	// We can perform DFS and memoize whether the snake has been in a certain
	// state before. It's not too bad.
	//
	// Note that since it is a grid and n >= 2, there is no edge case where the
	// snake is already at the end.
	//
	n := len(grid)
	seen := make(map[state]bool)
	st := func(i1, j1, i2, j2, dir int) state {
		return state{
			pos: [2][2]int{{i1, j1}, {i2, j2}},
			dir: uint8(dir),
		}
	}
	curr := []state{st(0, 0, 0, 1, dirRight)}
	next := []state{}
	ok := func(p [2]int) bool {
		i := p[0]
		j := p[1]
		return i >= 0 && j >= 0 && i < n && j < n && grid[i][j] != 1
	}
	movepos := func(p, d [2]int) [2]int {
		return [2]int{
			p[0] + d[0],
			p[1] + d[1],
		}
	}
	want := [2][2]int{{n - 1, n - 2}, {n - 1, n - 1}}
	seen[curr[0]] = true
	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			// this may need to change to include direction...
			// the problem description is not the best.
			if x.pos == want {
				return steps
			}

			// This will be messy, but so be it.
			// If facing right, we can try to rotate down
			if x.dir == dirRight {
				p1 := movepos(x.pos[0], [2]int{1, 0})
				p2 := movepos(x.pos[1], [2]int{1, 0})
				if ok(p1) && ok(p2) {
					t := x
					t.pos[1] = p1
					t.dir = dirDown
					if !seen[t] {
						seen[t] = true
						next = append(next, t)
					}
				}
			} else {
				// Facing down
				p1 := movepos(x.pos[0], [2]int{0, 1})
				p2 := movepos(x.pos[1], [2]int{0, 1})
				if ok(p1) && ok(p2) {
					t := x
					t.pos[1] = p1
					t.dir = dirRight
					if !seen[t] {
						seen[t] = true
						next = append(next, t)
					}
				}
			}

			// Move right
			r1 := movepos(x.pos[0], [2]int{0, 1})
			r2 := movepos(x.pos[1], [2]int{0, 1})
			if ok(r1) && ok(r2) {
				t := x
				t.pos[0] = r1
				t.pos[1] = r2
				if !seen[t] {
					seen[t] = true
					next = append(next, t)
				}
			}

			// Move down
			d1 := movepos(x.pos[0], [2]int{1, 0})
			d2 := movepos(x.pos[1], [2]int{1, 0})
			if ok(d1) && ok(d2) {
				t := x
				t.pos[0] = d1
				t.pos[1] = d2
				if !seen[t] {
					seen[t] = true
					next = append(next, t)
				}
			}
		}
		curr, next = next, curr
	}
	return -1
}

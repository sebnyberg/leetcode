package p3

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumMinutes(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,0,2,2,1,1,0,2,1,1,2,2,0,2,2,1,2,0,1,2,2,0,1,2,2,1,2,2],[2,2,2,1,1,2,2,1,2,0,1,1,1,2,2,1,1,0,2,2,2,0,1,0,1,2,2,2],[0,0,1,1,0,1,2,0,1,1,1,1,0,2,0,2,0,2,1,1,0,2,1,2,2,2,1,2],[2,2,0,0,0,0,1,0,1,0,2,0,1,0,2,0,0,1,2,1,0,1,1,1,2,0,2,0],[2,2,1,1,1,1,1,0,0,0,0,2,0,1,1,1,1,2,0,2,1,1,2,0,2,0,2,0],[0,1,0,1,2,2,2,0,2,0,2,2,1,2,0,0,1,0,2,0,2,0,1,2,2,0,2,0],[1,0,2,2,2,0,2,0,2,0,2,0,1,0,2,2,0,2,1,1,1,0,1,0,1,1,0,0],[0,1,2,0,1,0,1,0,2,1,2,0,1,1,1,1,0,1,1,0,0,2,0,1,0,1,0,2],[2,1,1,0,1,1,2,2,1,2,2,1,0,1,0,0,0,2,1,0,2,2,1,2,1,2,0,1],[1,1,2,0,2,2,1,2,0,2,1,1,0,0,0,2,2,2,2,1,2,2,0,2,1,1,2,0],[2,1,2,2,0,0,1,0,1,2,1,0,1,0,2,0,0,1,1,0,2,0,2,0,1,2,2,0],[1,0,1,1,0,0,0,0,0,1,0,2,0,2,1,2,1,1,0,1,0,0,2,1,2,1,0,2],[2,0,1,0,2,0,1,0,2,0,2,1,2,0,2,2,2,1,0,2,1,0,1,2,1,0,1,1],[0,2,2,1,0,2,1,0,1,2,2,1,2,2,1,2,0,1,2,2,0,2,1,0,2,1,0,0],[0,2,2,2,1,2,1,0,0,2,2,0,1,0,2,1,0,0,2,1,1,1,2,1,2,1,0,1],[2,2,2,1,1,1,1,0,2,2,2,1,0,0,2,2,0,0,1,1,0,0,2,1,2,1,2,2],[2,1,2,1,1,1,0,2,1,0,1,1,2,1,0,0,1,1,2,1,2,2,1,2,0,2,0,0]]"),
			-1,
		},
		{
			leetcode.ParseMatrix("[[0,2,1,1,0],[1,2,0,0,1],[2,2,1,1,0]]"),
			-1,
		},
		{
			leetcode.ParseMatrix("[[0,2,0,0,1],[0,2,0,2,2],[0,2,0,0,0],[0,0,2,2,0],[0,0,0,0,0]]"),
			0,
		},
		{
			leetcode.ParseMatrix("[[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[0,0,0,0],[0,1,2,0],[0,2,0,0]]"),
			-1,
		},
		{
			leetcode.ParseMatrix("[[0,0,0],[2,2,0],[1,2,0]]"),
			1000000000,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maximumMinutes(tc.grid))
		})
	}
}

const (
	fire   = 1 << 0
	wall   = 1 << 1
	person = 1 << 2
)

func maximumMinutes(grid [][]int) int {
	grid[0][0] |= person

	m, n := len(grid), len(grid[0])

	// Bounds / wall check
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j]&wall == 0
	}

	// Collect original fire positions
	origFires := [][2]int{}
	state := make([][]int, m)
	for i := range grid {
		state[i] = make([]int, n)
		for j := range grid[i] {
			if grid[i][j] == fire {
				origFires = append(origFires, [2]int{i, j})
			}
		}
	}

	currFire := make([][2]int, len(origFires))
	nextFire := make([][2]int, 0, len(origFires))
	currPerson := [][2]int{}
	nextPerson := [][2]int{}

	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	// Check whether a person can reach the end after waiting a number of turns
	survivesAfterTurns := func(turnsToWait int) bool {
		// Reset state
		for i := range grid {
			copy(state[i], grid[i])
		}
		currFire = currFire[:len(origFires)]
		currPerson = currPerson[:0]
		copy(currFire, origFires)

		// Perform BFS
		// If there is no place left for the fire or person to move, then the person
		// will die (by virtue of being stuck somewhere)
		// Note that t<=turnsToWait+1 is necessary because the person is only added
		// to `currPerson` after turnsToWait+1
		for t := 1; len(currFire) > 0 || len(currPerson) > 0 || t <= turnsToWait+1; t++ {
			nextPerson = nextPerson[:0]
			nextFire = nextFire[:0]
			if t == turnsToWait+1 {
				currPerson = append(currPerson, [2]int{0, 0})
			}

			// Move each person position one step further
			for _, p := range currPerson {
				// Note that the fire can cover the person's current position due to the
				// ordering of the current and next loop:
				if state[p[0]][p[1]]&fire > 0 {
					continue
				}
				for _, d := range dirs {
					i, j := p[0]+d[0], p[1]+d[1]
					if !ok(i, j) || state[i][j]&person > 0 || state[i][j]&fire > 0 {
						continue
					}
					if i == m-1 && j == n-1 { // Reached the end => person survives
						return true
					}
					state[i][j] |= person
					nextPerson = append(nextPerson, [2]int{i, j})
				}
			}

			// Move each fire one step further
			for _, f := range currFire {
				for _, d := range dirs {
					i, j := f[0]+d[0], f[1]+d[1]
					if !ok(i, j) || state[i][j]&fire > 0 {
						continue
					}
					if i == m-1 && j == n-1 { // Reached the end => person dies
						return false
					}
					state[i][j] |= fire
					nextFire = append(nextFire, [2]int{i, j})
				}
			}
			nextFire, currFire = currFire, nextFire
			nextPerson, currPerson = currPerson, nextPerson
		}
		return false
	}

	// Fast-track
	if !survivesAfterTurns(0) { // person will never survive
		return -1
	}
	if survivesAfterTurns(m * n) { // person cannot die
		return 1e9
	}

	// Perform binary search, finding the lowest number of turns which results in
	// the death of the person.
	lo, hi := 0, m*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if survivesAfterTurns(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

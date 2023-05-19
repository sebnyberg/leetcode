package p1263minimummovestomoveaboxtotheirtargetlocation

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPushBox(t *testing.T) {
	for i, tc := range []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				[]byte("######"),
				[]byte("#T####"),
				[]byte("#..B.#"),
				[]byte("#.##.#"),
				[]byte("#...S#"),
				[]byte("######"),
			},
			3,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minPushBox(tc.grid))
		})
	}
}

func minPushBox(grid [][]byte) int {
	// Just do exhaustive search. Not a very hard problem imo, just a lot of
	// code.
	type state struct {
		player [2]int
		box    [2]int
	}
	start := state{}
	var target [2]int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'B' {
				start.box = [2]int{i, j}
			}
			if grid[i][j] == 'S' {
				start.player = [2]int{i, j}
			}
			if grid[i][j] == 'T' {
				target = [2]int{i, j}
			}
		}
	}
	curr := []state{start}
	next := []state{}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != '#'
	}
	seen := make(map[state]int)
	seen[start] = 0
	res := math.MaxInt32
	if target == start.box {
		return 0
	}

	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x.player[0] + d[0]
				jj := x.player[1] + d[1]
				if !ok(ii, jj) {
					continue
				}
				nextState := x
				nextState.player = [2]int{ii, jj}
				// if next position is the box and the position two blocks
				// forward is empty, we may push the box to its next position
				currPushes := seen[x]
				if nextState.player == x.box {
					iii := ii + d[0]
					jjj := jj + d[1]
					if !ok(iii, jjj) {
						continue
					}
					// push box forward
					nextState.box = [2]int{iii, jjj}
					currPushes++
					if nextState.box == target {
						res = min(res, currPushes)
						continue
					}
				}
				v, exists := seen[nextState]
				if exists && v <= currPushes {
					continue
				}
				seen[nextState] = currPushes
				next = append(next, nextState)
			}
		}
		curr, next = next, curr
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

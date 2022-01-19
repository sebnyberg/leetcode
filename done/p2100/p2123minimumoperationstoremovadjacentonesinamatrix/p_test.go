package p2123minimumoperationstoremovadjacentonesinamatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumOperations(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 0}, {0, 0, 1, 0}}, 4},
		{[][]int{{1, 1, 1, 1, 1, 1, 1}}, 3},
		{[][]int{{1, 1, 0}, {0, 1, 1}, {1, 1, 1}}, 3},
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0},
		{[][]int{{0, 1}, {1, 0}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minimumOperations(tc.grid))
		})
	}
}

func minimumOperations(grid [][]int) int {
	// New attempt:
	//
	// Any subset of the grid containing un-isolated cells form a group that can be
	// considered a bipartite graph. While it is not required to flip exactly one
	// group in the bipartite graph - it would be ok to flip two ones to zeroes -
	// it is probably (?) optimal. Let's try and see if it works.

	// Turns out that is not optimal, instead we must parse an actual graph and
	// try to flip so that the result is minimal.
	n := len(grid[0])
	m := len(grid)
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	var minMoves int

	for i := range grid {
		for j, v := range grid[i] {
			if v == 0 {
				continue
			}
			graph := collectGraph(grid, seen, i, j, m, n)
			minMoves += min(len(graph[0]), len(graph[1]))
		}
	}

	return minMoves
}

type coord struct {
	i, j int
}

var near = [4]coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func collectGraph(grid [][]int, seen [][]bool, startI, startJ, m, n int) [2][]coord {
	cur := []coord{{startI, startJ}}
	next := []coord{}
	var res [2][]coord

	seen[startI][startJ] = true
	for insIdx := 0; len(cur) > 0; insIdx = (insIdx + 1) % 2 {
		next = next[:0]
		res[insIdx] = append(res[insIdx], cur...)
		// Expand to neighbours
		for _, p := range cur {
			for _, d := range near {
				ii, jj := p.i+d.i, p.j+d.j
				if ii < 0 || jj < 0 || ii >= m || jj >= n || grid[ii][jj] == 0 || seen[ii][jj] {
					continue
				}
				// Add to next
				next = append(next, coord{ii, jj})
				seen[ii][jj] = true
			}
		}
		cur, next = next, cur
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0317shortestdistancefromallbuildings

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestDistance(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}, 7},
		{[][]int{{1, 0}}, 1},
		{[][]int{{1}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, shortestDistance(tc.grid))
		})
	}
}

type pos struct {
	i, j int
}

func shortestDistance(grid [][]int) int {
	bfs, bfs2 := []pos{}, []pos{}
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range grid {
		dist[i] = make([]int, n)
	}
	mark := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell != 1 {
				continue
			}
			// BFS for this building
			bfs = bfs[:0]
			bfs = append(bfs, pos{i, j})
			level := 1
			for len(bfs) > 0 {
				bfs2 = bfs2[:0]
				for _, p := range bfs {
					for _, near := range []pos{
						{p.i - 1, p.j}, {p.i + 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1},
					} {
						if near.i < 0 || near.i >= m || near.j < 0 || near.j >= n || grid[near.i][near.j] != mark {
							continue
						}
						grid[near.i][near.j]--
						dist[near.i][near.j] += level
						bfs2 = append(bfs2, near)
					}
				}
				bfs, bfs2 = bfs2, bfs
				level++
			}
			mark--
		}
	}

	res := math.MaxInt32
	// Find the smallest distance place that has the final mark
	for i, row := range grid {
		for j, cell := range row {
			if cell == mark {
				res = min(res, dist[i][j])
			}
		}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func near(p pos, m, n int) []pos {
	res := make([]pos, 0)
	for _, p2 := range []pos{
		{p.i - 1, p.j}, {p.i + 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1},
	} {
		if p2.i < 0 || p2.i >= m || p2.j < 0 || p2.j >= m {
			continue
		}
		res = append(res, p2)
	}
	return res
}

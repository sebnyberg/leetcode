package amz2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestDistance(t *testing.T) {
	for _, tc := range []struct {
		maze        [][]int
		start       []int
		destination []int
		want        int
	}{
		{[][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{0, 4}, []int{3, 2}, -1},
		{[][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4}, 12},
		{[][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{4, 3}, []int{0, 1}, -1},
		{[][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{4, 3}, []int{0, 1}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.maze), func(t *testing.T) {
			require.Equal(t, tc.want, shortestDistance(tc.maze, tc.start, tc.destination))
		})
	}
}

type pos struct {
	i, j int
	dist int
}

func (p *pos) coord() [2]int {
	return [2]int{p.i, p.j}
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
	cur := []pos{{start[0], start[1], 0}}
	bfs2 := []pos{}
	seen := make(map[[2]int]int)
	seen[cur[0].coord()] = 0
	shortestDist := math.MaxInt32
	for len(cur) > 0 {
		bfs2 = bfs2[:0]
		for _, p := range cur {
			// For each direction the ball can be kicked toward
			for _, dir := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				// Kick the ball as long as possible
				p2 := p
				for p2.i >= 0 && p2.i < m && p2.j >= 0 && p2.j < n && maze[p2.i][p2.j] != 1 {
					p2.dist++
					p2.i += dir[0]
					p2.j += dir[1]
				}

				// Adjust out of bounds
				p2.dist--
				p2.i -= dir[0]
				p2.j -= dir[1]
				prev, exists := seen[p2.coord()]
				if exists && prev <= p2.dist {
					continue
				}
				if p2.i == destination[0] && p2.j == destination[1] {
					shortestDist = min(shortestDist, p2.dist)
				}
				seen[p2.coord()] = p2.dist
				bfs2 = append(bfs2, p2)
			}
		}
		cur, bfs2 = bfs2, cur
	}
	if shortestDist == math.MaxInt32 {
		return -1
	}
	return shortestDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

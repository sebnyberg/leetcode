package p1197minknightmoves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minKnightMoves(t *testing.T) {
	for _, tc := range []struct {
		x    int
		y    int
		want int
	}{
		{0, -300, 150},
		{2, 1, 1},
		{5, 5, 4},
	} {
		t.Run(fmt.Sprintf("%v,%v", tc.x, tc.y), func(t *testing.T) {
			require.Equal(t, tc.want, minKnightMoves(tc.x, tc.y))
		})
	}
}

func minKnightMoves(x int, y int) int {
	// Perform BFS, where valid moves either:
	//
	// 1. Reduces the mahattan distance, or
	// 2. Is within 5 manhattan distances from the target
	//
	// There is probably an exact solution as well, but I cannot
	// be arsed to try all combinations of positions you could
	// find yourself in
	//
	// Update: yup, there is an exact solution, but I would
	// still prefer this one since it does not require any
	// mental capacity at all - suits me well :)
	//
	var cur pos
	target := pos{x, y}
	bfs, bfs2 := []pos{cur}, []pos{}
	var visited [700][700]bool
	offset := 350
	visited[offset][offset] = true
	var moves int
	for len(bfs) > 0 {
		for _, p := range bfs {
			if p.x == x && p.y == y {
				return moves
			}
			d := dist(p, target)
			for _, near := range p.GetMoves() {
				if visited[offset+near.x][offset+near.y] {
					continue
				}
				visited[offset+near.x][offset+near.y] = true
				// Check if position is closer to the target than
				// then previous position OR less than 5 dist from the target
				newDist := dist(near, target)
				if newDist < d || newDist <= 5 {
					bfs2 = append(bfs2, near)
				}
			}
		}
		bfs, bfs2 = bfs2, bfs
		moves++
	}
	return 0
}

type pos struct {
	x, y int
}

func dist(p1, p2 pos) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p pos) GetMoves() []pos {
	return []pos{
		{p.x + 1, p.y + 2},
		{p.x + 2, p.y + 1},
		{p.x - 1, p.y + 2},
		{p.x - 2, p.y + 1},
		{p.x + 1, p.y - 2},
		{p.x + 2, p.y - 1},
		{p.x - 1, p.y - 2},
		{p.x - 2, p.y - 1},
	}
}

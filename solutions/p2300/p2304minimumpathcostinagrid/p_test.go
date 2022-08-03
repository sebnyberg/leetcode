package p2304minimumpathcostinagrid

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPathCost(t *testing.T) {
	for _, tc := range []struct {
		grid     [][]int
		moveCost [][]int
		want     int
	}{
		{
			leetcode.ParseMatrix("[[5,3],[4,0],[2,1]]"),
			leetcode.ParseMatrix("[[9,8],[1,5],[10,12],[18,6],[2,4],[14,3]]"),
			17,
		},
		{
			leetcode.ParseMatrix("[[5,1,2],[4,0,3]]"),
			leetcode.ParseMatrix("[[12,10,15],[20,23,8],[21,7,1],[8,1,13],[9,10,25],[5,3,2]]"),
			6,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minPathCost(tc.grid, tc.moveCost))
		})
	}
}

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	curr := make([]int, n)
	next := make([]int, n)
	for i := range curr {
		curr[i] = grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			next[j] = grid[i][j]
			minCost := math.MaxInt32
			for k := 0; k < n; k++ {
				cellValue := grid[i-1][k]
				cost := moveCost[cellValue][j]
				minCost = min(minCost, curr[k]+cost)
			}
			next[j] += minCost
		}
		curr, next = next, curr
	}
	minCost := math.MaxInt32
	for _, v := range curr {
		minCost = min(minCost, v)
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

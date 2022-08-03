package p0573squirrelsimulation

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	for _, tc := range []struct {
		height   int
		width    int
		tree     []int
		squirrel []int
		nuts     [][]int
		want     int
	}{
		{5, 7, []int{2, 2}, []int{4, 4}, [][]int{{3, 0}, {2, 5}}, 12},
		{1, 3, []int{0, 1}, []int{0, 0}, [][]int{{0, 2}}, 3},
		{
			5,
			5,
			[]int{3, 2},
			[]int{0, 1},
			leetcode.ParseMatrix("[[2,0],[4,1],[0,4],[1,3],[1,0],[3,4],[3,0],[2,3],[0,2],[0,0],[2,2],[4,2],[3,3],[4,4],[4,0],[4,3],[3,1],[2,1],[1,4],[2,4]]"),
			100,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.height), func(t *testing.T) {
			require.Equal(t, tc.want, minDistance(tc.height, tc.width, tc.tree, tc.squirrel, tc.nuts))
		})
	}
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	// Find closest nut to the squirrel
	dist := func(a, b []int) int {
		dx := abs(a[0] - b[0])
		dy := abs(a[1] - b[1])
		return dx + dy
	}

	var baseMoves int
	for _, nut := range nuts {
		baseMoves += dist(nut, tree) * 2
	}

	minMoves := math.MaxInt32
	for _, nut := range nuts {
		regularDist := dist(nut, tree) * 2
		firstNutDist := dist(squirrel, nut) + dist(nut, tree)
		minMoves = min(minMoves, baseMoves-regularDist+firstNutDist)
	}

	return minMoves
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

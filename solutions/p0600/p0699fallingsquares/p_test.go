package p0699fallingsquares

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fallingSquares(t *testing.T) {
	for _, tc := range []struct {
		positions [][]int
		want      []int
	}{
		{
			leetcode.ParseMatrix("[[2,1],[2,9],[1,8]]"),
			[]int{1, 10, 18},
		},
		{
			leetcode.ParseMatrix("[[1,2],[2,3],[6,1]]"),
			[]int{2, 5, 5},
		},
		{
			leetcode.ParseMatrix("[[100,100],[200,100]]"),
			[]int{100, 100},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.positions), func(t *testing.T) {
			require.Equal(t, tc.want, fallingSquares(tc.positions))
		})
	}
}

func fallingSquares(positions [][]int) []int {
	// Use coordinate compression and half-open intervals
	// Below is boilerplate which maps each x-position to a compressed index.
	xsMap := make(map[int]struct{})
	for _, p := range positions {
		xsMap[p[0]] = struct{}{}
		xsMap[p[0]+p[1]] = struct{}{}
	}
	xs := make([]int, 0, len(xsMap))
	for x := range xsMap {
		xs = append(xs, x)
	}
	sort.Ints(xs)
	xIdx := make(map[int]int)
	for i, x := range xs {
		xIdx[x] = i
	}

	heights := make([]int, len(xs))
	var res []int
	var maxHeightTotal int
	for _, p := range positions {
		start, end := xIdx[p[0]], xIdx[p[0]+p[1]]
		h := p[1]

		// Find highest box that intersects with this one (if any)
		var maxHeight int
		for x := xIdx[p[0]]; x < xIdx[p[0]+p[1]]; x++ {
			if heights[x] > maxHeight {
				maxHeight = heights[x]
				h = p[1] + heights[x]
			}
		}

		// Update heights
		for x := start; x < end; x++ {
			heights[x] = h
		}

		// Register max value so far
		maxHeightTotal = max(maxHeightTotal, h)
		res = append(res, maxHeightTotal)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

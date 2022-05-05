package p0598rangeaddition2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCount(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{
			92, 2,
			[][]int{{70, 1}, {37, 1}, {3, 2}, {67, 1}, {37, 2}, {87, 2}, {26, 1}, {43, 1}, {19, 1}, {63, 1}, {67, 1}, {19, 1}, {14, 2}, {5, 1}, {27, 2}, {44, 2}, {13, 1}},
			3,
		},
		{3, 3, [][]int{
			{2, 2},
			{3, 3}, {3, 3}, {3, 3},
			{2, 2},
			{3, 3}, {3, 3}, {3, 3},
			{2, 2},
			{3, 3}, {3, 3}, {3, 3},
		}, 4},
		{3, 3, [][]int{}, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, maxCount(tc.m, tc.n, tc.ops))
		})
	}
}

func maxCount(m int, n int, ops [][]int) int {
	// Register ops one by one
	// The final region will be the minimum x and y values from ops
	minX, minY := n, m
	var count int
	for _, op := range ops {
		minX = min(minX, op[1])
		minY = min(minY, op[0])
		count++
	}
	return minX * minY
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

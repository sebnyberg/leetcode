package p1465maxareaofpieceofcakeafterhorizontalandverticalcuts

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxArea(t *testing.T) {
	for _, tc := range []struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
		want           int
	}{
		{5, 4, []int{3}, []int{3}, 9},
		{5, 4, []int{3, 1}, []int{1}, 6},
		{5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.h), func(t *testing.T) {
			require.Equal(t, tc.want, maxArea(tc.h, tc.w, tc.horizontalCuts, tc.verticalCuts))
		})
	}
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	prev := 0
	maxH := 0
	for _, cut := range horizontalCuts {
		maxH = max(maxH, cut-prev)
		prev = cut
	}
	maxH = max(maxH, h-prev)

	sort.Ints(verticalCuts)
	prev = 0
	maxW := 0
	for _, cut := range verticalCuts {
		maxW = max(maxW, cut-prev)
		prev = cut
	}
	maxW = max(maxW, w-prev)

	return (maxH * maxW) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

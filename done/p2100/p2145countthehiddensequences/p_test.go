package p2145countthehiddensequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfArrays(t *testing.T) {
	for _, tc := range []struct {
		differences  []int
		lower, upper int
		want         int
	}{
		{[]int{-40}, -46, 53, 60},
		{[]int{1, -3, 4}, 1, 6, 2},
		{[]int{3, -4, 5, 1, -2}, -4, 5, 4},
		{[]int{4, -7, 2}, 3, 6, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.differences), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfArrays(tc.differences, tc.lower, tc.upper))
		})
	}
}

func numberOfArrays(differences []int, lower int, upper int) int {
	minVal := 0
	maxVal := 0
	var val int
	for _, d := range differences {
		val += d
		minVal = min(minVal, val)
		maxVal = max(maxVal, val)
	}
	maxDelta := upper - lower + 1
	actualDelta := maxVal - minVal + 1
	if actualDelta > maxDelta {
		return 0
	}
	return maxDelta - actualDelta + 1
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

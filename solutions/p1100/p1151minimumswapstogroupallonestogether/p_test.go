package p1151minimumswapstogroupallonestogether

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwaps(t *testing.T) {
	for _, tc := range []struct {
		data []int
		want int
	}{
		{
			[]int{1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
			30,
		},
		{[]int{1, 0, 1, 0, 1}, 1},
		{[]int{0, 0, 0, 1, 0}, 0},
		{[]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.data), func(t *testing.T) {
			require.Equal(t, tc.want, minSwaps(tc.data))
		})
	}
}

func minSwaps(data []int) int {
	n := len(data)
	var windowSize int
	for _, d := range data {
		if d == 1 {
			windowSize++
		}
	}

	var l, r int
	var requiredSwaps int
	for r = 0; r < windowSize; r++ {
		if data[r] == 0 {
			requiredSwaps++
		}
	}
	minSwaps := requiredSwaps
	for r < n {
		// Remove left number in window
		if data[l] == 0 {
			requiredSwaps--
		}
		l++
		// Add right number in window
		if data[r] == 0 {
			requiredSwaps++
		}
		r++
		if requiredSwaps < minSwaps {
			minSwaps = requiredSwaps
		}
	}
	return minSwaps
}

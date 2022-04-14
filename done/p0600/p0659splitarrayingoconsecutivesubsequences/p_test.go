package p0659splitarrayingoconsecutivesubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPossible(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 4, 5}, false},
		{[]int{1, 2, 3, 5, 5, 6, 7}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 3, 4, 5}, true},
		{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, isPossible(tc.nums))
		})
	}
}

func isPossible(nums []int) bool {
	var counts [2004]int16
	const offset = 1001
	for _, v := range nums {
		counts[v+offset]++
	}
	start := []int{}
	for i, count := range counts {
		n := int16(len(start))
		// Check longest sequences (to be removed)
		for j := int16(0); j < n-count; j++ {
			if i-start[j] < 3 {
				return false
			}
		}
		// Append current index for each new sequence
		for j := int16(0); j < count-n; j++ {
			start = append(start, i)
		}
		start = start[max(0, n-count):]
	}

	return true
}

func max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

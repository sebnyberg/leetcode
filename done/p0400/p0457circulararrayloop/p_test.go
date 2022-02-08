package p0457circulararrayloop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_circularArrayLoop(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{5, -1, 1, 2, 2}, false},
		{[]int{1, 1, 2}, true},
		{[]int{-1, -2, -3, -4, -5}, false},
		{[]int{1, 1}, true},
		{[]int{2, -1, 1, 2, 2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2, 1, -1, -2, -2}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, circularArrayLoop(tc.nums))
		})
	}
}

func circularArrayLoop(nums []int) bool {
	// Due to pidgeon-hole principle, everything in nums is a cycle, the question
	// is whether any of those cycles contain all positive or all negative entries
	const intSize = 32 << (^uint(0) >> 63)

	n := len(nums)
	var empty [5001]bool
	hasUniDirCycle := func(start int) bool {
		if (start+nums[start]+n)%n == start {
			return false
		}
		seen := empty
		var neg int
		var pos int
		i := start
		for {
			if seen[i] {
				return i == start && neg^pos == 1
			}
			seen[i] = true
			neg |= -(nums[i] >> (intSize - 1))
			pos |= int(-uint(nums[i]) >> (intSize - 1))
			i = (i + nums[i] + n*1000) % n
		}
	}
	for i := 0; i < n-1; i++ {
		if hasUniDirCycle(i) {
			return true
		}
	}
	return false
}

package p2444countsubarraywithfixedbounds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubarrays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		minK int
		maxK int
		want int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countSubarrays(tc.nums, tc.minK, tc.maxK))
		})
	}
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	maxPos := -1
	minPos := -1
	startPos := -1
	var res int64
	for i, x := range nums {
		if x < minK || x > maxK {
			maxPos = i
			minPos = i
			startPos = i
			continue
		}
		if x == minK {
			minPos = i
		}
		if x == maxK {
			maxPos = i
		}
		dist := int64(min(maxPos, minPos) - startPos)
		res += dist
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

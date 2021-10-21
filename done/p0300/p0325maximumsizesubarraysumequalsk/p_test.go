package p0325maximumsizesubarraysumequalsk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubArrayLen(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, -1, 5, -2, 3}, 3, 4},
		{[]int{-2, -1, 2, 1}, 1, 2},
		{[]int{-2, -1, 1, 3, -3, 4}, 5, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSubArrayLen(tc.nums, tc.k))
		})
	}
}

func maxSubArrayLen(nums []int, k int) int {
	sum := 0
	seenAt := make(map[int]int)
	seenAt[0] = -1
	var maxLen int
	for i, n := range nums {
		sum += n
		if j, exists := seenAt[sum-k]; exists {
			maxLen = max(maxLen, i-j)
		}
		if _, exists := seenAt[sum]; !exists {
			seenAt[sum] = i
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1708largestsubarraylengthk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 4, 5, 2, 3}, 3, []int{5, 2, 3}},
		{[]int{1, 4, 5, 2, 3}, 4, []int{4, 5, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, largestSubarray(tc.nums, tc.k))
		})
	}
}

func largestSubarray(nums []int, k int) []int {
	var maxval int
	var maxidx int
	n := len(nums)
	for i := 0; i <= n-k; i++ {
		if nums[i] > maxval {
			maxval = nums[i]
			maxidx = i
		}
	}
	return nums[maxidx : maxidx+k]
}

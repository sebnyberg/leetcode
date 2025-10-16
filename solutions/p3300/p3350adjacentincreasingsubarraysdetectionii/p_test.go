package p3350adjacentincreasingsubarraysdetectionii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxIncreasingSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxIncreasingSubarrays(tc.nums))
		})
	}
}

func maxIncreasingSubarrays(nums []int) int {
	var count int
	var counts []int
	for i := range nums {
		if i > 0 && nums[i] <= nums[i-1] {
			counts = append(counts, count)
			count = 0
		}
		count++
	}
	counts = append(counts, count)
	ok := func(k int) bool {
		for i := range counts {
			if counts[i] >= k*2 || i > 0 && counts[i-1] >= k && counts[i] >= k {
				return true
			}
		}
		return false
	}
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ok(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

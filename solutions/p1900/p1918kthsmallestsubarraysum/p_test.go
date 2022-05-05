package p1918kthsmallestsubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallestSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 3}, 4, 3},
		{[]int{3, 3, 5, 5}, 7, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallestSubarraySum(tc.nums, tc.k))
		})
	}
}

func kthSmallestSubarraySum(nums []int, k int) int {
	var maxSum int
	for _, n := range nums {
		maxSum += n
	}
	l, r := 0, maxSum
	for l < r {
		mid := l + (r-l)/2
		if countLeqSum(nums, mid) < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func countLeqSum(nums []int, maxSum int) int {
	var count int
	sum := 0
	l := 0
	for r := 0; r < len(nums); {
		if sum+nums[r] <= maxSum {
			sum += nums[r]
			count += r - l + 1
			r++
		} else {
			sum -= nums[l]
			l++
		}
	}
	return count
}

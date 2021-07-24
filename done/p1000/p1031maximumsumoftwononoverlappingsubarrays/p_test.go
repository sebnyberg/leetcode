package p1031maximumsumoftwononoverlappingsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumTwoNoOverlap(t *testing.T) {
	for _, tc := range []struct {
		nums                []int
		firstLen, secondLen int
		want                int
	}{
		{[]int{8, 20, 6, 2, 20, 17, 6, 3, 20, 8, 12}, 5, 4, 108},
		{[]int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15}, 3, 5, 109},
		{[]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3, 31},
		{[]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2, 20},
		{[]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2, 29},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumTwoNoOverlap(tc.nums, tc.firstLen, tc.secondLen))
		})
	}
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	var sum int

	// Calculate left sum
	left := make([]int, len(nums)+1)
	for i := 0; i < firstLen; i++ {
		sum += nums[i]
	}
	left[firstLen] = sum
	for i := firstLen; i < n; i++ {
		sum += nums[i]
		sum -= nums[i-firstLen]
		left[i+1] = max(left[i], sum)
	}

	// Calculate right sum
	right := make([]int, len(nums)+1)
	sum = 0
	for i := n - 1; i >= n-firstLen; i-- {
		sum += nums[i]
	}
	right[n-firstLen] = sum
	for i := n - firstLen - 1; i >= 0; i-- {
		sum += nums[i]
		sum -= nums[i+firstLen]
		right[i] = max(right[i+1], sum)
	}

	// For each position for the second window, calculate the max possible sum
	sum = 0
	res := 0
	for i := 0; i < secondLen; i++ {
		sum += nums[i]
	}
	res = max(res, sum+max(left[0], right[secondLen]))
	for i := secondLen; i < n; i++ {
		sum += nums[i]
		sum -= nums[i-secondLen]
		res = max(res, sum+max(left[i-secondLen+1], right[i+1]))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

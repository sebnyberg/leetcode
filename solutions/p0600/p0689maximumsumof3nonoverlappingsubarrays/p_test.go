package p0689maximumsumof3nonoverlappingsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2, []int{0, 2, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumOfThreeSubarrays(tc.nums, tc.k))
		})
	}
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	windowedLen := n - k + 1
	windowSum := make([]int, windowedLen)
	currSum := 0
	for i := 0; i < n; i++ {
		currSum += nums[i]
		if i >= k {
			currSum -= nums[i-k]
		}
		if i >= k-1 {
			windowSum[i-k+1] = currSum
		}
	}

	// Left contains the index for which a k-width interval's sum is maximized
	left := make([]int, windowedLen)
	var maxIdx int
	for i := 0; i < windowedLen; i++ {
		if windowSum[i] > windowSum[maxIdx] {
			maxIdx = i
		}
		left[i] = maxIdx
	}

	// Right contains the index for which a k-width interval's sum is maximized
	right := make([]int, windowedLen)
	maxIdx = windowedLen - 1
	for i := windowedLen - 1; i >= 0; i-- {
		if windowSum[i] >= windowSum[maxIdx] {
			maxIdx = i
		}
		right[i] = maxIdx
	}

	// Fix center position and find set of positions which yield the highest
	// total sum.
	res := make([]int, 3)
	for i := range res {
		res[i] = -1
	}
	maxSum := 0
	for j := k; j < windowedLen-k; j++ {
		l := left[j-k]
		r := right[j+k]
		cand := windowSum[l] + windowSum[j] + windowSum[r]
		if res[0] == -1 || cand > maxSum {
			maxSum = cand
			res[0] = l
			res[1] = j
			res[2] = r
		}
	}
	return res
}

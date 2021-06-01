package p1695maximumerasevalue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumUniqueSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{10000}, 10000},
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumUniqueSubarray(tc.nums))
		})
	}
}

func maximumUniqueSubarray(nums []int) int {
	// Two pointers: l, r
	l, r := 0, 1
	var counts [10001]int
	counts[nums[l]]++

	// While counts[nums[r]] == 0, move r
	// While nums[l] != nums[r], move l
	n := len(nums)
	sum := nums[l]
	maxSum := sum
	for r != n {
		for r < n && counts[nums[r]] == 0 {
			counts[nums[r]]++
			sum += nums[r]
			r++
		}
		// nums[r] now points to a number that would cause a duplicate
		maxSum = max(maxSum, sum)
		if r == n {
			break
		}

		// Introduce nums[r] to the set
		sum += nums[r]
		counts[nums[r]]++
		r++

		// Move left pointer until counts[nums[r]] is equal to one again
		for counts[nums[r-1]] != 1 {
			sum -= nums[l]
			counts[nums[l]]--
			l++
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

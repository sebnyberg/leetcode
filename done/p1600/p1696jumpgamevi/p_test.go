package p1696jumpgamevi

import (
	"container/list"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxResult(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, -1, -2, 4, -7, 3}, 2, 7},
		{[]int{10, -5, -2, 4, 0, 3}, 3, 17},
		{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxResult(tc.nums, tc.k))
		})
	}
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	// Slide a window over the nums array, keeping the max value from the window
	dp[0] = nums[0]

	// Deque contains the indices of elements, front has the highest value
	deque := list.New()
	deque.PushBack(0)
	for i := 1; i < n; i++ {
		// Front in deque is the max value from the range
		// Pop items with indices smaller than the current
		for deque.Len() > 0 && deque.Front().Value.(int) < i-k {
			deque.Remove(deque.Front())
		}

		// Max value is max of window + current
		dp[i] = dp[deque.Front().Value.(int)] + nums[i]

		// Add the current value by popping smaller values from the left
		// Any smaller value at the back of the deque is "lower quality" than the
		// current since they will fall out of favour faster than the current.
		for deque.Len() > 0 && dp[deque.Back().Value.(int)] <= dp[i] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)
	}
	return dp[n-1]
}

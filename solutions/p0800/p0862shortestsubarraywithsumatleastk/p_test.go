package p0852shortestsubarraywithsumatleastk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 4, -1},
		{[]int{2, -1, 2}, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, shortestSubarray(tc.nums, tc.k))
		})
	}
}

func shortestSubarray(nums []int, k int) int {
	// Let's consider positive values only. The easiest way to find the best
	// result is a two-pointer / stack approach. Add elements within a window
	// until its sum is >= k. Then evaluate candidates while moving the left
	// pointer and decrementing the sum. Then move right pointer, etc.
	//
	// For negative values, it is a bit more difficult. We do not know how far we
	// should move either pointer, because moving the pointer is not monotonically
	// increasing or decreasing the sum.
	//
	// To solve this problem, we use a double-ended queue. Each index in the queue
	// is a possible starting position for a range with a positive sum (w.r.t. the
	// current position).
	//
	// If the sum between the current index and the first index in the deque is
	// >= k, then that must be the optimal solution that starts with that first
	// element. Hence, we can poll the element and continue.
	//
	// If the sum between the current index and the last index in the queue is
	// negative, then that index can never be relevant for finding a solution.
	// This is because any solution must contain a positive sum. Hence, we pop
	// elements from the end until the range has a positive sum (or there are no
	// elements left).
	n := len(nums)
	presum := make([]int, n+1)
	for i, n := range nums {
		presum[i+1] = presum[i] + n
	}
	deque := make([]int, 0)
	res := n + 1
	for i := 0; i <= n; i++ {
		// Any range within deque that is >= k is a valid answer.
		// We can dequeue because any later index must have a larger range than the
		// current one.
		for len(deque) > 0 && presum[i]-presum[deque[0]] >= k {
			res = min(res, i-deque[0])
			deque = deque[1:]
		}

		// The remove indices from the deque until the range between its last index
		// and the current has a positive sum. Otherwise, it would make no sense to
		// evaluate the range (it would yield a sum <= 0)
		for len(deque) > 0 && presum[deque[len(deque)-1]]-presum[i] <= 0 {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}
	if res <= n {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0862shortestsubarraywithsumatleastk

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
		{[]int{-34, 37, 51, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6}, 151, 2},
		{[]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453, 9},
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
	n := len(nums)
	presum := make([]int, n+1)
	for i, n := range nums {
		presum[i+1] = presum[i] + n
	}
	deque := make([]int, 0)
	res := n + 1
	for i := 0; i <= n; i++ {
		for len(deque) > 0 && presum[i]-presum[deque[0]] >= k {
			res = min(res, i-deque[0])
			deque = deque[1:]
		}
		for len(deque) > 0 && presum[i] <= presum[deque[len(deque)-1]] {
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

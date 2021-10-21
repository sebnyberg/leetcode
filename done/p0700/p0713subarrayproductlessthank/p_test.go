package p0713subarrayproductlessthank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numSubarrayProductLessThanK(tc.nums, tc.k))
		})
	}
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	// Second try
	cur := 1
	start := 0
	count := 0
	for i, n := range nums {
		cur *= n
		for start < i && cur >= k {
			cur /= nums[start]
			start++
		}
		if cur < k {
			count += i - start + 1
		}
	}
	return count
}

// First try
func numSubarrayProductLessThanKStack(nums []int, k int) int {
	// Idea: keep a stack of numbers below k
	// For each element in nums, multiply it with current elements in the stack,
	// removing any elements greater than or equal to k, then add the size of the
	// stack to the result.
	n := len(nums)
	stack := make([]int, 0, n/2)
	var count int
	for _, n := range nums {
		var offset int
		for i, x := range stack {
			stack[i] = x * n
		}
		stack = append(stack, n)
		for _, x := range stack {
			if x >= k {
				offset++
			} else {
				break
			}
		}
		stack = stack[offset:]
		count += len(stack)
	}
	return count
}

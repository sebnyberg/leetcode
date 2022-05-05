package p1950maximumofminimumvaluesinallsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaximums(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 2, 4}, []int{4, 2, 1, 0}},
		{[]int{10, 20, 50, 10}, []int{50, 20, 10, 10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaximums(tc.nums))
		})
	}
}

func findMaximums(nums []int) []int {
	// Create a monotonic increasing stack.
	// Read numbers from nums.
	// Pop numbers from the stack until the encountered number is larger than
	// all items in the stack.
	// A popped item is guaranteed to be the smallest item between the prior
	// index and the index where that item was encountered, this is the
	// length of the window for which to update the maximum minimum value.
	n := len(nums)
	nums = append(nums, 0) // sentinel
	var stack intStack
	maxVals := make([]int, n)
	for i, num := range nums {
		for len(stack) > 0 && nums[stack.peek()] >= num {
			back := nums[stack.pop()]
			d := i - 1
			if len(stack) > 0 {
				d -= stack.peek() + 1
			}
			maxVals[d] = max(maxVals[d], back)
		}
		stack.push(i)
	}
	for i := n - 1; i > 0; i-- {
		maxVals[i-1] = max(maxVals[i-1], maxVals[i])
	}
	return maxVals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type intStack []int

func (s *intStack) push(x int) {
	*s = append(*s, x)
}
func (s intStack) peek() int {
	return (s)[len(s)-1]
}
func (s *intStack) pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

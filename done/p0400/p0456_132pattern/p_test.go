package p0456_132pattern

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_find132pattern(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{3, 1, 4, 2}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{-1, 3, 2, 0}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, find132pattern(tc.nums))
		})
	}
}

type IntStack []int

func (s *IntStack) Pop() int {
	n := len(*s)
	item := (*s)[n-1]
	(*s) = (*s)[:n-1]
	return item
}

func (s *IntStack) Peek() int {
	n := len(*s)
	return (*s)[n-1]
}

func (s *IntStack) Push(x int) {
	(*s) = append((*s), x)
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n <= 2 {
		return false
	}

	// Given any position "j" > 0 in nums,
	// a pattern of 132 is most likely to happen for the "i" which holds
	// the smallest number before "j".
	// minLeft[j] contains the smallest possible number before j.
	minLeft := make([]int, n)
	minLeft[0] = math.MaxInt32
	for i := 1; i < n; i++ {
		minLeft[i] = min(minLeft[i-1], nums[i-1])
	}

	// For any position j, we can test against minLeft[j] to verify the 13 pattern.
	// If minLeft[j] < nums[j], we need to verify the right side.
	//
	// For any given k > j, there are three cases:
	// 1. nums[k] <= minLeft[j]. Since minLeft[j] = min(nums[:j-1]),
	//		there is no way for nums[k] to ever be part of a 132. Pop from stack.
	// 2. nums[k] >= nums[j]. There is no 132 pattern, but nums[k] could be part
	//		of a pattern for a lower j. Add it to the stack
	// 3. nums[k] < nums[j] and nums[j] > minLeft[j]: SUCCESS.
	//
	// IMPORTANT!
	// Right stack is sorted. This is due to the fact that if it were
	// unsorted, the 132 pattern would have happened already.
	rightGreaterThanCur := make(IntStack, 0)
	rightGreaterThanCur.Push(nums[n-1])
	for j := n - 2; j > 0; j-- {
		if nums[j] <= minLeft[j] {
			continue
		}

		// Case 1. minLeft will never become smaller, so items can be safely removed.
		for len(rightGreaterThanCur) > 0 && rightGreaterThanCur.Peek() <= minLeft[j] {
			rightGreaterThanCur.Pop()
		}

		// 132 pattern success!
		if len(rightGreaterThanCur) > 0 && rightGreaterThanCur.Peek() < nums[j] {
			return true
		}

		// Case 2. Store number for future use (nums[j] may be unusually small)
		rightGreaterThanCur.Push(nums[j])
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

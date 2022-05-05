package p0503nextgreaterelementii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElements(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, nextGreaterElements(tc.nums))
		})
	}
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums2 := make([]int, n*2)
	copy(nums2, nums)
	copy(nums2[n:], nums)
	res := make([]int, n*2)
	for i := range res {
		res[i] = -1
	}
	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		num := nums2[i]
		// Pop elements from stack
		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {
			// Add value to result
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res[:n]
}

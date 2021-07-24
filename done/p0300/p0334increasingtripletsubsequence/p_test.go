package p0334increasingtripletsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_increasingTriplet(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 1, 1, 1}, false},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, increasingTriplet(tc.nums))
		})
	}
}

func increasingTriplet(nums []int) bool {
	// Patience sort O(n) time complexity, O(3) space (exit cond on 3)
	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		for j := len(stack); j >= 0; j-- {
			if j == 0 {
				stack[0] = nums[i]
				break
			}
			if stack[j-1] > nums[i] {
				continue
			}
			if stack[j-1] == nums[i] {
				break
			}
			if j >= len(stack) {
				stack = append(stack, nums[i])
			} else {
				stack[j] = nums[i]
			}
			break
		}
		if len(stack) >= 3 {
			return true
		}
	}
	return false
}

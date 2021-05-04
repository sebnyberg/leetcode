package p0665nondecreasingarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPossibility(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 4, 8, 3, 4}, false},
		{[]int{1, 4, 2, 3}, true},
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, checkPossibility(tc.nums))
		})
	}
}

func checkPossibility(nums []int) bool {
	for i, n := range nums {
		if i == 0 {
			continue
		}
		if n < nums[i-1] {
			// Either the current or the previous number is invalid
			// Attempt to fix both and check result
			newPrev := -10001
			if i > 1 {
				newPrev = nums[i-2]
			}
			prev := nums[i-1]
			nums[i-1] = newPrev
			if checkOK(nums[i-1:]) {
				return true
			}
			nums[i-1] = prev
			nums[i] = nums[i-1]
			return checkOK(nums[i:])
		}
	}

	return true
}

func checkOK(nums []int) bool {
	for i, n := range nums {
		if i == 0 {
			continue
		}
		if n < nums[i-1] {
			return false
		}
	}
	return true
}

package p2369checkifthereisavalidpartitionforthearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validPartition(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{4, 4, 4, 5, 6}, true},
		{[]int{1, 1, 1, 2}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, validPartition(tc.nums))
		})
	}
}

func validPartition(nums []int) bool {
	// Let's try DP and see if it fails
	// For any position, we may choose to jump either two or three steps forward
	// The next two or three values must match the criteria.
	// We may use DP to check whether we've visited a certain step before.
	// If something is seen, then it's invalid
	// We short-circuit on a solution
	seen := make([]bool, len(nums))
	res := dp(seen, nums, 0)
	return res
}

func dp(seen []bool, nums []int, i int) bool {
	if i == len(nums) {
		return true
	}
	if seen[i] {
		return false
	}
	seen[i] = true
	// Try all three cases
	if i+2 > len(nums) {
		return false
	}
	if nums[i] == nums[i+1] && dp(seen, nums, i+2) {
		return true
	}
	if i+3 > len(nums) {
		return false
	}
	a := nums[i] == nums[i+1] && nums[i+1] == nums[i+2]
	b := nums[i+1] == nums[i]+1 && nums[i+2] == nums[i]+2
	if (a || b) && dp(seen, nums, i+3) {
		return true
	}
	return false
}

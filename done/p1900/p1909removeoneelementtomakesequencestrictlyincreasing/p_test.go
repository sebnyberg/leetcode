package p1909removeoneelementtomakesequencestrictlyincreasing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canBeIncreasing(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 10, 5, 7}, true},
		{[]int{100, 21, 100}, true},
		{[]int{2, 3, 1, 2}, false},
		{[]int{1, 1, 1}, false},
		{[]int{1, 2, 3}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, canBeIncreasing(tc.nums))
		})
	}
}

func canBeIncreasing(nums []int) bool {
	var didRemove bool
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if didRemove {
				return false
			}
			didRemove = true
			// Go with current number
			if i <= 1 || nums[i-2] <= nums[i] {
				continue
			}
			// Go with previous
			nums[i] = nums[i-1]
		}
	}
	return true
}

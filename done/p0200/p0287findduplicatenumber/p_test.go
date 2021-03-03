package p0287findduplicatenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findDuplicate(tc.nums))
		})
	}
}

func findDuplicate(nums []int) int {
	// Floyd's hare and tortoise cycle detection
	slow, fast := nums[nums[0]], nums[nums[nums[0]]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

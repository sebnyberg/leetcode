package p1827minoperationstomakearrayincreasing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 5, 2, 4, 1}, 14},
		{[]int{8}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums))
		})
	}
}

func minOperations(nums []int) int {
	n := len(nums)
	var operations int
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			operations += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return operations
}

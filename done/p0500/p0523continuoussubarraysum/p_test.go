package p0523continuoussubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 23, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, checkSubarraySum(tc.nums, tc.k))
		})
	}
}

func checkSubarraySum(nums []int, k int) bool {
	modSum := nums[0] % k
	seen := make(map[int]bool)
	prev := modSum
	for _, n := range nums[1:] {
		modSum += n
		modSum %= k
		if modSum == 0 || seen[modSum] {
			return true
		}
		seen[prev] = true
		prev = modSum
	}
	return false
}

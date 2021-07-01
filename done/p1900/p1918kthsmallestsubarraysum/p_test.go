package p1918kthsmallestsubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallestSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 3}, 4, 3},
		{[]int{3, 3, 5, 5}, 7, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallestSubarraySum(tc.nums, tc.k))
		})
	}
}

func kthSmallestSubarraySum(nums []int, k int) int {
	return -1
}

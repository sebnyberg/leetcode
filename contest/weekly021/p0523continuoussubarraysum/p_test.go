package p0523continuoussubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkSubarraySum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{0, 0}, 7, true},
		{[]int{23, 2, 4, 6, 6}, 7, true},
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 13, false},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, checkSubarraySum(tc.nums, tc.k))
		})
	}
}

func checkSubarraySum(nums []int, k int) bool {
	sums := make(map[int]struct{}, len(nums))

	var cur int
	var prevSum int
	for _, num := range nums {
		cur += num
		want := cur % k
		if _, exists := sums[want]; exists {
			return true
		}
		sums[prevSum] = struct{}{}
		prevSum = cur % k
	}

	return false
}

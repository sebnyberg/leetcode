package p0560subarraysumequalsk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subarraySum(tc.nums, tc.k))
		})
	}
}

func subarraySum(nums []int, k int) int {
	sums := make(map[uint32]uint32, len(nums))
	sums[0] = 1
	var res uint32
	var sum uint32
	for _, n := range nums {
		sum += uint32(n)
		if v, exists := sums[sum-uint32(k)]; exists {
			res += v
		}
		sums[sum]++
	}
	return int(res)
}

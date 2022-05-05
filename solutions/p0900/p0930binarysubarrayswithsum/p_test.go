package p0930binarysubarrayswithsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubarraysWithSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		goal int
		want int
	}{
		{[]int{1, 0, 1, 0, 1}, 2, 4},
		{[]int{0, 0, 0, 0, 0}, 0, 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numSubarraysWithSum(tc.nums, tc.goal))
		})
	}
}

func numSubarraysWithSum(nums []int, goal int) int {
	preCount := make([]int, len(nums)+1)
	preCount[0] = 1
	var count int
	var sum int
	for _, n := range nums {
		sum += n
		if sum >= goal {
			count += preCount[sum-goal]
		}
		preCount[sum]++
	}
	return count
}

package p0494targetsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTargetSumWays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1, 256},
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findTargetSumWays(tc.nums, tc.target))
		})
	}
}

func findTargetSumWays(nums []int, target int) int {
	prev := make(map[int]int, len(nums))
	curr := make(map[int]int, len(nums))
	prev[nums[0]]++
	prev[-nums[0]]++
	for _, num1 := range nums[1:] {
		for k := range curr {
			delete(curr, k)
		}
		for s, count := range prev {
			curr[s+num1] += count
			curr[s-num1] += count
		}
		curr, prev = prev, curr
	}
	return prev[target]
}

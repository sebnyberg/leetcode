package p2770maximumnumberofjumbstoreachthelastindex

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumJumps(t *testing.T) {
	for i, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 6, 4, 1, 2}, 2, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumJumps(tc.nums, tc.target))
		})
	}
}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = math.MinInt32
	}
	res[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if abs(nums[j]-nums[i]) <= target {
				res[i] = max(res[i], res[j]+1)
			}
		}
	}
	if res[n-1] < 0 {
		return -1
	}
	return res[n-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

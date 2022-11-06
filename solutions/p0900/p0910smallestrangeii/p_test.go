package p0910smallestrangeii

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestRangeII(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{0, 10}, 2, 6},
		{[]int{1}, 0, 0},
		{[]int{1, 3, 6}, 3, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, smallestRangeII(tc.nums, tc.k))
		})
	}
}

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return 0
	}
	n := len(nums)
	d := nums[n-1] - nums[0]
	for i := 1; i < n; i++ {
		a := min(nums[0]+k, nums[i]-k)
		b := max(nums[n-1]-k, nums[i-1]+k)
		d = min(d, b-a)
	}
	return d
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

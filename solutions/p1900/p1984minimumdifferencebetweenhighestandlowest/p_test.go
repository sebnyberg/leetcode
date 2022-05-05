package p1984minimumdifferencebetweenhighestandlowest

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{90}, 1, 0},
		{[]int{9, 4, 1, 7}, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDifference(tc.nums, tc.k))
		})
	}
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := k - 1; i < len(nums); i++ {
		minDiff = min(minDiff, nums[i]-nums[i-k+1])
	}
	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

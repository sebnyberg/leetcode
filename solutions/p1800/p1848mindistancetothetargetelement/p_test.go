package p1848mindistancetothetargetelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDist(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		start  int
		want   int
	}{
		{[]int{5, 3, 6}, 5, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 5, 3, 1},
		{[]int{1}, 1, 0, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, getMinDistance(tc.nums, tc.target, tc.start))
		})
	}
}

func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if start+i < n && nums[start+i] == target {
			return i
		}
		if start-i >= 0 && nums[start-i] == target {
			return i
		}
	}
	return 0
}

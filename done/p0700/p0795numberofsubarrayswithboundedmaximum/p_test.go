package p0795numberofsubarrayswithboundedmaximum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubarrayBoundedMax(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		left  int
		right int
		want  int
	}{
		{[]int{2, 1, 4, 3}, 2, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numSubarrayBoundedMax(tc.nums, tc.left, tc.right))
		})
	}
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	n := len(nums)
	count := 0
	for l := 0; l < n; l++ {
		var maxVal int
		for r := l; r < n; r++ {
			maxVal = max(maxVal, nums[r])
			if maxVal > right {
				break
			}
			if maxVal >= left {
				count++
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

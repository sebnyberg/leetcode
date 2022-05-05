package p0162findpeakelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPeakElement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findPeakElement(tc.nums))
		})
	}
}

func findPeakElement(nums []int) int {
	n := len(nums)
	for i := range nums {
		if i > 0 && nums[i-1] >= nums[i] {
			continue
		}
		if i < n-1 && nums[i+1] >= nums[i] {
			continue
		}
		return i
	}
	return -1
}

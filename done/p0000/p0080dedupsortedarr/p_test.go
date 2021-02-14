package p0080dedupsortedarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		want     int
		wantNums []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			got := removeDuplicates(tc.nums)
			require.Equal(t, tc.want, got)
			tc.nums = tc.nums[:got]
			require.Equal(t, tc.wantNums, tc.nums)
		})
	}
}

func removeDuplicates(nums []int) int {
	i := 2
	for {
		if i >= len(nums) {
			break
		}
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
			continue
		}
		i++
	}
	return len(nums)
}

package p2210counthillsandvalleysinanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countHillValley(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 4, 1, 1, 6, 5}, 3},
		{[]int{6, 6, 5, 5, 4, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countHillValley(tc.nums))
		})
	}
}

func countHillValley(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	nums = nums[:j+1]
	var count int
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
			count++
		} else if nums[i-1] > nums[i] && nums[i+1] > nums[i] {
			count++
		}
	}
	return count
}

package p2104sumofsubarrayranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subArrayRanges(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{1, 2, 3}, 4},
		{[]int{1, 3, 3}, 4},
		{[]int{4, -2, -3, 4, 1}, 59},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subArrayRanges(tc.nums))
		})
	}
}

func subArrayRanges(nums []int) int64 {
	var sum int
	for i := 0; i < len(nums)-1; i++ {
		min, max := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}
			sum += max - min
		}
	}
	return int64(sum)
}

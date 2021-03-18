package p0376wigglesubseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wiggleMaxLength(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 0}, 1},
		{[]int{0, 0, 1, 1, 1}, 2},
		{[]int{84}, 1},
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{33, 53, 12, 64, 50, 41, 45, 21, 97, 35, 47, 92, 39, 0, 93, 55, 40, 46, 69, 42, 6, 95, 51, 68, 72, 9, 32, 84, 34, 64, 6, 2, 26, 98, 3, 43, 30, 60, 3, 68, 82, 9, 97, 19, 27, 98, 99, 4, 30, 96, 37, 9, 78, 43, 64, 4, 65, 30, 84, 90, 87, 64, 18, 50, 60, 1, 40, 32, 48, 50, 76, 100, 57, 29, 63, 53, 46, 57, 93, 98, 42, 80, 82, 9, 41, 55, 69, 84, 82, 79, 30, 79, 18, 97, 67, 23, 52, 38, 74, 15}, 67},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, wiggleMaxLength(tc.nums))
		})
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Find first "wiggle"
	var i int
	for i = 1; i < len(nums) && nums[i] == nums[i-1]; i++ {
	}
	if i == len(nums) {
		return 1
	}

	curLen := 2
	pos := nums[i] > nums[i-1]
	prev := nums[i]
	for ; i < len(nums); i++ {
		switch {
		case pos && nums[i] > prev:
			prev = nums[i]
			continue
		case !pos && nums[i] < prev:
			prev = nums[i]
			continue
		case nums[i] == prev:
			continue
		default:
			pos = !pos
			prev = nums[i]
			curLen++
		}
	}

	return curLen
}

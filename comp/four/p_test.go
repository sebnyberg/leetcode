package one_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMiddleIndex(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -1, 8, 4}, 3},
		{[]int{-1000, 1000}, -1},
		{[]int{1, -1, 4}, 2},
		{[]int{2, 5}, -1},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMiddleIndex(tc.nums))
		})
	}
}

func findMiddleIndex(nums []int) int {
	var rightSum int
	for _, n := range nums {
		rightSum += n
	}
	var leftSum int
	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

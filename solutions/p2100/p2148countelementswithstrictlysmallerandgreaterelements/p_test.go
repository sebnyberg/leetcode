package p2148countelementswithstrictlysmallerandgreaterelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countElements(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{11, 7, 2, 15}, 2},
		{[]int{-3, 3, 3, 90}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countElements(tc.nums))
		})
	}
}

func countElements(nums []int) int {
	min, max := nums[0], nums[0]
	for _, x := range nums {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	var count int
	for _, x := range nums {
		if x > min && x < max {
			count++
		}
	}
	return count
}

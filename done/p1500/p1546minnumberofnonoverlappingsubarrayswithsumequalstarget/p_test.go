package p1546minnumberofnonoverlappingsubarrayswithsumequalstarget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNonOverlapping(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 2, 2},
		{[]int{-1, 3, 5, 1, 4, 2, -9}, 6, 2},
		{[]int{-2, 6, 6, 3, 5, 4, 1, 2, 8}, 10, 3},
		{[]int{0, 0, 0}, 0, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxNonOverlapping(tc.nums, tc.target))
		})
	}
}

func maxNonOverlapping(nums []int, target int) int {
	seen := make(map[int]struct{})
	seen[0] = struct{}{}
	presum := 0
	var res int
	for _, num := range nums {
		presum += num
		if _, exists := seen[presum-target]; exists {
			res++
			seen = make(map[int]struct{})
			seen[0] = struct{}{}
			presum = 0
			continue
		}
		seen[presum] = struct{}{}
	}
	return res
}

package p0001twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tcs := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			ij := twoSum(tc.nums, tc.target)
			i, j := ij[0], ij[1]
			assert.Equal(t, tc.want[0], i, "i")
			assert.Equal(t, tc.want[1], j, "j")
		})
	}
}

func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums[i+1:] {
			if n1+n2 == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{0, 0}
}

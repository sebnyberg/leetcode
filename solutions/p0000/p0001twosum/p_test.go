package p0001twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
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
			require.Equal(t, tc.want, twoSum(tc.nums, tc.target))
		})
	}
}

func twoSum(nums []int, target int) []int {
	numIdx := make(map[int]int, len(nums))
	for i, n1 := range nums {
		if j, exists := numIdx[target-n1]; exists {
			return []int{j, i}
		}
		numIdx[n1] = i
	}
	return nil
}

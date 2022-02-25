package p0525contiguousarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxLength(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{0, 0, 0, 1, 1, 1}, 6},
		{[]int{0, 0, 0, 0, 1, 1}, 4},
		{[]int{0, 0, 1, 0, 0, 0, 1, 1}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxLength(tc.nums))
		})
	}
}

func findMaxLength(nums []int) int {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	n := len(nums)
	var sum int
	var res int
	firstSumIdx := make(map[int]int, n)
	firstSumIdx[0] = -1
	for i, num := range nums {
		sum += num
		if j, exists := firstSumIdx[sum]; exists {
			res = max(res, i-j)
		} else {
			firstSumIdx[sum] = i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

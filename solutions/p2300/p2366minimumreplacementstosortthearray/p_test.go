package p2366minimumreplacementstosortthearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumReplacement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{7, 6, 15, 6, 11, 14, 10}, 10},
		{[]int{12, 9, 7, 6}, 6},
		{[]int{12, 9, 7, 6, 17, 19, 21}, 6},
		{[]int{3, 9, 3}, 2},
		{[]int{2, 10, 20, 19, 1}, 47},
		{[]int{1, 2, 3, 4, 5}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumReplacement(tc.nums))
		})
	}
}

func minimumReplacement(nums []int) int64 {
	var ops int64
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			continue
		}
		rest := nums[i] % nums[i+1]
		parts := nums[i] / nums[i+1]
		if rest > 0 {
			parts++
		}
		ops += int64(parts - 1)
		nums[i] /= parts
	}
	return ops
}

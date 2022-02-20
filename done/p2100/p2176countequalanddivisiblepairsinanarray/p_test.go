package p2176

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 1, 2, 2, 2, 1, 3}, 2, 4},
		{[]int{1, 2, 3, 4}, 1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countPairs(tc.nums, tc.k))
		})
	}
}

func countPairs(nums []int, k int) int {
	var res int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				res++
			}
		}
	}
	return res
}

package p2216minimumdeletionstomakearraybeautiful

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDeletion(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2, 3, 5}, 1},
		{[]int{1, 1, 2, 2, 3, 3}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minDeletion(tc.nums))
		})
	}
}

func minDeletion(nums []int) int {
	// Length must be even
	// nums[i] != nums[i+1] for all i%2 == 0
	// I think this can be done in a greedy fashion...

	// j is the current index in the result slice,
	// i is the current index in the original slice
	var count int
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || count&1 == 1 || nums[i] != nums[i+1] {
			count++
		}
	}
	if count%2 == 1 {
		count--
	}
	res := len(nums) - count
	return res
}

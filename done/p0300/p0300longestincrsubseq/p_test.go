package p0300longestincrsubseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLIS(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLIS(tc.nums))
		})
	}
}

func lengthOfLIS(nums []int) int {
	// Approach: Patience Sort - O(n)
	stack := make([]int, 1, 10)
	stack[0] = -10001 // Dummy guard
	for _, num := range nums {
		var pos int
		for pos = len(stack); stack[pos-1] >= num; pos-- {
		}
		if pos == len(stack) {
			stack = append(stack, num)
		} else {
			stack[pos] = num
		}
	}
	return len(stack) - 1
}

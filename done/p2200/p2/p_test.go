package p2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_triangularSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 8},
		{[]int{5}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, triangularSum(tc.nums))
		})
	}
}

func triangularSum(nums []int) int {
	curr := nums
	next := make([]int, len(nums))
	n := len(nums)
	for ; n > 1; n-- {
		for i := 0; i < n-1; i++ {
			next[i] = (curr[i] + curr[i+1]) % 10
		}
		curr, next = next, curr
	}
	return curr[0]
}

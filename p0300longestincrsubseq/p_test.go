package p0300longestincrsubseq

import (
	"fmt"
	"sort"
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
	dp := make([]int, 0, len(nums))
	for _, num := range nums {
		insertPos := sort.SearchInts(dp, num)
		if insertPos == len(dp) {
			dp = append(dp, num)
		} else {
			dp[insertPos] = num
		}
	}
	return len(dp)
}

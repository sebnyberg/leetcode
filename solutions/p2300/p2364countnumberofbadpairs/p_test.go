package p2364countnumberofbadpairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBadPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{4, 1, 3, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countBadPairs(tc.nums))
		})
	}
}

func countBadPairs(nums []int) int64 {
	count := make(map[int]int)
	var res int64
	for i := range nums {
		adj := nums[i] - i
		totalCount := i - count[adj]
		res += int64(totalCount)
		count[adj]++
	}
	return res
}

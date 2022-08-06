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

// https://www.github.com/sebnyberg/leetcode
func countBadPairs(nums []int) int64 {
	// The idea here is to adjust each number so that its position is subtracted
	// from its value.
	// That way numbers from the same sequence have the same number.
	// Counting bad pairs is then a matter of counting non-equal prior numbers.
	count := make(map[int]int, len(nums)/10)
	var res int64
	for i := range nums {
		x := nums[i] - i
		totalCount := i - count[x]
		res += int64(totalCount)
		count[x]++
	}
	return res
}

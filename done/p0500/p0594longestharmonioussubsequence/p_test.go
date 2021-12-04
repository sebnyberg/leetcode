package p0594longestharmonioussubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLHS(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 1, 1, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findLHS(tc.nums))
		})
	}
}

func findLHS(nums []int) int {
	valCount := make(map[int]uint32)
	for _, num := range nums {
		valCount[num]++
	}
	var maxCount uint32
	for val, count := range valCount {
		if c := valCount[val-1]; c > 0 && c+count > maxCount {
			maxCount = c + count
		}
	}
	return int(maxCount)
}

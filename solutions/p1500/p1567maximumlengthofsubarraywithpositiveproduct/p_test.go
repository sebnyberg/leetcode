package p1567maximumlengthofsubarraywithpositiveproduct

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaxLen(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, -2, -3, 4}, 4},
		{[]int{0, 1, -2, -3, -4}, 3},
		{[]int{-1, -2, -3, 0, 1}, 2},
		{[]int{-1, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, getMaxLen(tc.nums))
		})
	}
}

func getMaxLen(nums []int) int {
	// Any section which includes a zero cannot be an answer to the problem.
	// For any other section, it makes sense to keep a positive and negative sum
	// with their corresponding lengths, and when finding a zero, pick the
	// longest positive sum
	nums = append(nums, 0)
	res := 0
	var negCount int
	first, last := -1, 0
	var start int
	for i, num := range nums {
		if num < 0 {
			negCount++
			if first == -1 {
				first = i
			}
			last = i
			continue
		}
		if num == 0 {
			if negCount%2 == 1 {
				res = max(res, max(last-start, i-first-1))
			} else {
				res = max(res, i-start)
			}
			negCount = 0
			first, last = -1, 0
			start = i + 1
			continue
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

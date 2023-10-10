package p2786visitarraypositionstomaximizescore

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxScore(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		x    int
		want int64
	}{
		{[]int{2, 3, 6, 1, 9, 2}, 5, 13},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxScore(tc.nums, tc.x))
		})
	}
}

func maxScore(nums []int, x int) int64 {
	var maxOdd int
	var maxEven int
	if nums[0]%2 == 0 {
		maxOdd = math.MinInt32
	} else {
		maxEven = math.MinInt32
	}
	for i := range nums {
		if nums[i]%2 == 0 {
			maxEven = max(maxEven, nums[i]+maxEven)
			if maxOdd >= 0 {
				maxEven = max(maxEven, nums[i]+maxOdd-x)
			}
		} else {
			maxOdd = max(maxOdd, nums[i]+maxOdd)
			if maxEven >= 0 {
				maxOdd = max(maxOdd, nums[i]+maxEven-x)
			}
		}
	}
	return int64(max(maxEven, maxOdd))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

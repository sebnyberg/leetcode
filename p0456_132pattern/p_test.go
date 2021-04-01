package p0456_132pattern

import (
	"fmt"
	"testing"
)

func Test_find132pattern(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		// {[]int{1, 2, 3, 4}, false},
		// {[]int{3, 1, 4, 2}, true},
		// {[]int{-1, 3, 2, 0}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			// require.Equal(t, tc.want, find132pattern(tc.nums))
		})
	}
}

// func find132pattern(nums []int) bool {
// 	minLeft, maxLeft := math.MaxInt32, math.MinInt32
// 	for _, n := range nums {
// 		if n < maxLeft && n > minLeft {
// 			return true
// 		}
// 		if
// 	}
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

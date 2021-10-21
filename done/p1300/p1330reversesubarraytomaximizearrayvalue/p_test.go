package p1330reversesubarraytomaximizearrayvalue

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxValueAfterReverse(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 5, 4}, 10},
		{[]int{2, 4, 9, 24, 2, 1, 10}, 68},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxValueAfterReverse(tc.nums))
		})
	}
}

func maxValueAfterReverse(nums []int) int {
	// Insane exercise... I just copied this from the discussion
	// For each combination of pairs (a,b) (c,d) where b != c,
	// min(a,b) could be either a, or b,
	// and max(a,b) would be the other
	// Similarly, min(c,d) could be either c or d, etc.
	// In order, this could result in 4*3*2*1 = 12 different permutations of
	// a,b,c,d
	// The only time swapping is guaranteed to improve the result is when
	// both min(a,b) and max(a,b) is smaller than min(c,d) and max(c,d), i.e.
	// when a,b <= c,d
	n := len(nums)
	// There are also edge cases where b == 0 and c == n-1
	total := 0
	res := 0
	smallestMaxOfPair := math.MaxInt32
	largestMinOfPair := math.MinInt32
	for i := 0; i < n-1; i++ {
		a, b := nums[i], nums[i+1]
		d := abs(a - b)
		total += d
		res = max(res, abs(nums[0]-b)-abs(a-b))   // swap current with first
		res = max(res, abs(a-nums[n-1])-abs(a-b)) // swap current with last
		smallestMaxOfPair = min(smallestMaxOfPair, max(a, b))
		largestMinOfPair = max(largestMinOfPair, min(a, b))
	}
	res = max(res, (largestMinOfPair-smallestMaxOfPair)*2)
	return total + res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package p0801minimumswapstomakesequencesincreasing

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwap(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			//   0   1  2  3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19
			//       1  1  2   2   2   3   3   3   4   4   4   4   4   8   5
			[]int{2, 1, 6, 9, 10, 13, 13, 16, 19, 26, 23, 24, 25, 27, 32, 31, 35, 36, 37, 39},
			[]int{0, 5, 8, 8, 10, 12, 14, 15, 22, 22, 28, 29, 30, 31, 30, 33, 33, 36, 37, 38},
			//       1  2  2   3   3   3   4   4   4   5   6   7   8   5   9
			6,
		},
		{[]int{0, 4, 4, 5, 9}, []int{0, 1, 6, 8, 10}, 1},
		{[]int{1, 3, 5, 4}, []int{1, 2, 3, 7}, 1},
		{[]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minSwap(tc.nums1, tc.nums2))
		})
	}
}

func minSwap(nums1 []int, nums2 []int) int {
	// Calculate the min cost for a pair being swapped and not swapped for a given
	// position.
	n := len(nums1)
	var curr [2]int
	curr[1] = 1
	for i := 1; i < n; i++ {
		next := [2]int{math.MaxInt32, math.MaxInt32}
		// First consider the current pair not being swapped and valid
		d11 := nums1[i] - nums1[i-1]
		d12 := nums1[i] - nums2[i-1]
		d21 := nums2[i] - nums1[i-1]
		d22 := nums2[i] - nums2[i-1]
		if d11 > 0 && d22 > 0 {
			next[0] = min(next[0], curr[0])
			next[1] = min(next[1], curr[1]+1)
		}
		if d12 > 0 && d21 > 0 {
			next[1] = min(next[1], curr[0]+1)
			next[0] = min(next[0], curr[1])
		}
		// If cross-deltas are invalid, then we must swap from a prior, valid,
		// swapped state
		if d12 <= 0 || d21 <= 0 {
			next[1] = min(next[1], curr[1]+1)
		}
		// If deltas are invalid, then we must swap from a prior valid unswapped
		// state.
		if d11 <= 0 || d22 <= 0 {
			next[1] = min(next[1], curr[0]+1)
		}
		curr = next
	}
	return min(curr[0], curr[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

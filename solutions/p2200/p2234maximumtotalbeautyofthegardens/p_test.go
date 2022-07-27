package p2234maximumtotalbeautyofthegardens

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumBeauty(t *testing.T) {
	for _, tc := range []struct {
		flowers    []int
		newFlowers int64
		target     int
		full       int
		partial    int
		want       int64
	}{
		{[]int{5, 19, 1, 1, 6, 10, 18, 12, 20, 10, 11}, 6, 20, 3, 11, 47},
		{[]int{13}, 18, 15, 9, 2, 28},
		{
			[]int{36131, 31254, 5607, 11553, 82824, 59230, 40967, 69571, 36874, 38700, 81107, 28500, 61796, 54371, 23405, 51780, 75265, 37257, 86314, 32258, 47254, 76690, 18014, 21538, 96700, 15094, 57253, 57073, 7284, 24501, 21412, 69582, 15724, 43829, 81444, 78281, 88953, 6671, 94646, 31037, 89465, 86033, 27431, 30774, 48592, 26067},
			2304903454,
			48476,
			5937,
			15214,
			737765815,
		},
		{[]int{20, 17, 1, 14, 8, 18}, 48, 20, 3, 4, 91},
		{[]int{18, 16, 10, 10, 5}, 10, 3, 15, 4, 75},
		{[]int{1, 3, 1, 1}, 7, 6, 12, 1, 14},
		{[]int{2, 4, 5, 3}, 10, 5, 2, 6, 30},
	} {
		t.Run(fmt.Sprintf("%+v", tc.flowers), func(t *testing.T) {
			require.Equal(t, tc.want, maximumBeauty(tc.flowers, tc.newFlowers, tc.target, tc.full, tc.partial))
		})
	}
}

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	// The solution must be O(n) or O(nlogn), so we're looking for some kind of
	// clever DP-type solution, greedy, or binary search.

	// Sorting most likely makes the solution easier to find.
	sort.Ints(flowers)
	n := len(flowers)

	// Fast-lane: if all flowers are already above target, return full*n
	allAbove := func() bool {
		for _, x := range flowers {
			if x < target {
				return false
			}
		}
		return true
	}
	if allAbove() {
		return int64(full * n)
	}

	// Preprocess input to calculate the cost of making all prior gardens have the
	// same amount of flowers as the current one.
	minValCost := make([]int, n+1)
	minValCost[0] = 0
	for i := 1; i < n; i++ {
		minValCost[i] = minValCost[i-1] + i*(flowers[i]-flowers[i-1])
	}
	minValCost[n] = math.MaxInt64

	var maxRes int

	// Fill zero to all gardens (or as far as it's possible),
	// calculating the maximum minimum value that is possible given that amount of
	// complete gardens using binary search.
	for i := n; i >= 0; i-- {
		if i < n && flowers[i] <= target {
			newFlowers -= int64(target - flowers[i])
		}
		if newFlowers < 0 {
			break
		}
		if i == 0 {
			maxRes = max(maxRes, full*n)
			break
		}
		// Find max min value
		l, r := 0, i
		for l < r {
			mid := l + (r-l)/2
			if minValCost[mid] > int(newFlowers) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		partialVal := flowers[l-1]
		if int(newFlowers) > minValCost[l-1] {
			// We may be able to go higher.
			d := int(newFlowers) - minValCost[l-1]
			d /= l
			partialVal = min(target-1, flowers[l-1]+d)
		}
		maxRes = max(maxRes, partialVal*partial+(n-i)*full)
	}

	return int64(maxRes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

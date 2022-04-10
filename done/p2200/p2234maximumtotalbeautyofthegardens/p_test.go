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
		{
			[]int{64753, 70721, 35179, 14240, 85904, 41870, 19694, 36688, 99646, 49085, 85619, 26435, 18771, 55583, 60427, 72027, 44868, 95581, 98990, 81882, 51628, 82625, 43425, 74385, 63862, 53800, 76884, 13139, 63703, 68373, 81076, 18220, 91633, 52737, 65764, 89790, 42570, 19317, 43152, 65395, 28911, 46582, 25554, 76178, 9173, 25436, 5842, 54867, 52028, 58952, 63094, 29583, 11598, 69236, 68720, 19564, 2757, 57855, 60757, 23659, 67839, 32038, 65734, 15290, 42556, 48913, 13312, 96397, 27608, 42084, 98068, 68397, 29079, 4376, 97930, 52427, 11075, 13088, 81742, 29956, 21925, 49807, 67952, 16287, 68255, 57174, 59840, 77891, 40221, 3156, 5128, 3497, 65228, 93320, 68648, 33111, 9676, 97523, 39295, 52427, 74753, 7798, 89393, 60840, 55642, 9150, 5029, 12613, 85348, 84827, 25601, 95020, 86573, 73326, 63948, 49911, 62048, 42886, 80491, 83517, 85604, 37006, 56101, 21357, 71697, 73275, 82457, 91209, 59471, 92914, 70091, 36917, 63223, 12256, 60120, 20354, 61169, 49300, 40539, 97203, 84401, 54030, 506, 83883, 75217, 63499, 46032, 7437, 96915, 32796, 98461, 14572, 11146, 5326, 17821, 86744, 44579, 44213, 66468, 20471, 38854, 44697, 99750, 63414, 24457, 66770, 62214, 60029, 18358, 87031, 11454, 96612, 31108, 8877, 27205, 40854, 90185, 52276, 59947, 67499, 16375, 33772, 5300, 591, 89451, 197, 73076, 74487, 40029, 77991, 71909, 19864, 27554, 64674, 27391, 88798, 92353, 54954, 39830, 57000, 65183, 79551, 47415, 70699, 53812, 79880, 96880, 55432, 85877, 67140, 16297, 90637, 581, 22493, 20043, 92338, 7151, 48283, 42379, 69667, 34839, 65754, 942, 7666, 30914, 47012, 27900, 19866, 98873, 86437, 38098, 57219, 24538, 34548, 6735, 86260, 98999, 42451, 34032, 18392, 26693, 92306, 18429, 86527, 19493, 56794, 27966, 79551, 22829, 99628, 5979, 37492, 83251, 40886, 23767, 72893, 79000, 29487, 74422, 39852},
			3463894,
			65243,
			40780,
			47,
			9945865,
		},
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

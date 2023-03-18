package p2594minimumtimetorepaircars

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_repairCars(t *testing.T) {
	for i, tc := range []struct {
		ranks []int
		cars  int
		want  int64
	}{
		{[]int{4, 3, 2, 1}, 10, 16},
		{[]int{5, 1, 8}, 6, 16},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, repairCars(tc.ranks, tc.cars))
		})
	}
}

func repairCars(ranks []int, cars int) int64 {
	// Typical binary search problem
	check := func(t int) bool {
		m := cars
		for _, r := range ranks {
			// r*n^2 <= t <==> n <= sqrt(t/r)
			n := math.Sqrt(float64(t / r))
			m -= int(n)
			if m <= 0 {
				return true
			}
		}
		return m <= 0
	}

	lo, hi := 0, math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int64(lo)
}

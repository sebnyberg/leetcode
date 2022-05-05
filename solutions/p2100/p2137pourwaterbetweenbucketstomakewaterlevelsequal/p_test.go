package p2137pourwaterbetweenbucketstomakewaterlevelsequal

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equalizeWater(t *testing.T) {
	for _, tc := range []struct {
		buckets []int
		loss    int
		want    float64
	}{
		{[]int{1, 2, 7}, 80, 2},
		{[]int{2, 4, 6}, 50, 3.5},
		{[]int{3, 3, 3, 3}, 40, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buckets), func(t *testing.T) {
			require.InEpsilon(t, tc.want, equalizeWater(tc.buckets, tc.loss), 1e-5)
		})
	}
}

func equalizeWater(buckets []int, loss int) float64 {
	sort.Ints(buckets)
	n := len(buckets)
	floatBuckets := make([]float64, n)
	for i := range buckets {
		floatBuckets[i] = float64(buckets[i])
	}
	keep := float64(100-loss) / 100

	canHoldOrMore := func(amt float64) bool {
		var want, got float64
		for _, b := range floatBuckets {
			if b < amt {
				want += amt - b
			} else {
				got += keep * (b - amt)
			}
		}
		return got >= want
	}

	var lo, hi float64 = 0, 100000
	for math.Abs(hi-lo) > 1e-5 {
		mid := (lo + hi) / 2
		if canHoldOrMore(mid) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

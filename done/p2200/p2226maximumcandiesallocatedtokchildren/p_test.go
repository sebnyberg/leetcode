package p2226maximumcandiesallocatedtokchildren

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumCandies(t *testing.T) {
	for _, tc := range []struct {
		candies []int
		k       int64
		want    int
	}{
		{[]int{5, 8, 6}, 3, 5},
		{[]int{2, 5}, 11, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.candies), func(t *testing.T) {
			require.Equal(t, tc.want, maximumCandies(tc.candies, tc.k))
		})
	}
}

func maximumCandies(candies []int, k int64) int {
	f := func(x int) bool {
		kk := k
		for _, c := range candies {
			kk -= int64(c / x)
			if kk <= 0 {
				break
			}
		}
		return kk <= 0
	}

	lo, hi := 1, math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		// Check if mid is possible
		if f(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

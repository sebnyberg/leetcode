package p2517maximumtastinessofcandybasket

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumTastiness(t *testing.T) {
	for i, tc := range []struct {
		price []int
		k     int
		want  int
	}{
		{[]int{13, 5, 1, 8, 21, 2}, 3, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumTastiness(tc.price, tc.k))
		})
	}
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	check := func(x int) bool {
		// Find at least k-1 jumps with stride x
		prev := price[0]
		kk := k - 1
		for i := range price {
			if price[i]-prev >= x {
				kk--
				if kk == 0 {
					return true
				}
				prev = price[i]
			}
		}
		return false
	}
	lo := 1
	hi := math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo == 0 {
		return 0
	}
	return lo - 1
}

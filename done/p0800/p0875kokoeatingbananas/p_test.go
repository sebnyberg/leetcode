package p0875kokoeatingbananas

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minEatingSpeed(t *testing.T) {
	for _, tc := range []struct {
		piles []int
		h     int
		want  int
	}{
		{[]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184},
			823855818, 14},
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
	} {
		t.Run(fmt.Sprintf("%+v", tc.piles), func(t *testing.T) {
			require.Equal(t, tc.want, minEatingSpeed(tc.piles, tc.h))
		})
	}
}

func minEatingSpeed(piles []int, h int) int {
	var max int
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	lo, hi := 1, max
	for lo < hi {
		mid := (lo + hi) / 2

		hh := h
		ok := true
		for _, p := range piles {
			hh -= ((p - 1) / mid) + 1
			if hh < 0 {
				ok = false
				break
			}
		}

		if ok {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

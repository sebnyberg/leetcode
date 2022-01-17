package p2141maximumrunningtimeofncomputers

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxRunTime(t *testing.T) {
	for _, tc := range []struct {
		n         int
		batteries []int
		want      int64
	}{
		{2, []int{3, 3, 3}, 4},
		{2, []int{1, 1, 1, 1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxRunTime(tc.n, tc.batteries))
		})
	}
}

func maxRunTime(n int, batteries []int) int64 {
	// It is 'bad' to have a battery with high power, and so the highest power
	// battery should always be plugged into a computer.

	// Since a battery's capacity can be 1e9, we must work with large chunks of
	// battery at a time.

	// Imagine len(batteries) == n, then all batteries must be plugged in non-stop
	// and the solution is min(batteries[:])

	// Imagine len(batteries) == n+1, then the highest n batteries will be plugged
	// in until the unused battery is larger in capacity than the lowest plugged
	// in one.

	// So an unused 'low' battery will eventually start offloading high batteries,
	// increasing their capacity.

	// So you could imagine shifting the capacity of a low battery to a high one,
	m := len(batteries)
	sort.Ints(batteries)
	var extra int
	for i := 0; i < m-n; i++ {
		extra += batteries[i]
	}
	batteries = batteries[m-n:]

	ok := func(level, extra int) bool {
		for i := range batteries {
			if batteries[i] < level {
				extra -= level - batteries[i]
			}
			if extra < 0 {
				return false
			}
		}
		return extra >= 0
	}

	lo, hi := 0, math.MaxInt64
	for lo < hi {
		mid := (lo + hi) / 2
		if ok(mid, extra) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int64(hi - 1)
}

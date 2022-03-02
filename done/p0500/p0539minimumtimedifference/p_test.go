package p0539minimumtimedifference

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinDifference(t *testing.T) {
	for _, tc := range []struct {
		timePoints []string
		want       int
	}{
		{[]string{"23:59", "00:00"}, 1},
		{[]string{"00:00", "23:59", "00:00"}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.timePoints), func(t *testing.T) {
			require.Equal(t, tc.want, findMinDifference(tc.timePoints))
		})
	}
}

func findMinDifference(timePoints []string) int {
	ts := make([]int, 0, 2*len(timePoints))
	for _, tp := range timePoints {
		var h, m int
		fmt.Sscanf(tp, "%02d:%02d", &h, &m)
		ts = append(ts, h*60+m, (h*60+m)+24*60)
	}
	sort.Ints(ts)
	minDist := 24 * 60
	for i := 1; i < len(ts); i++ {
		minDist = min(minDist, ts[i]-ts[i-1])
	}
	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

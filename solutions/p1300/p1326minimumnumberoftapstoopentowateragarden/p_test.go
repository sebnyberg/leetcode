package p1326minimumnumberoftapstoopentowateragarden

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minTaps(t *testing.T) {
	for i, tc := range []struct {
		n      int
		ranges []int
		want   int
	}{
		{5, []int{3, 4, 1, 1, 0, 0}, 1},
		{3, []int{0, 0, 0, 0}, -1},
		{35, []int{1, 0, 4, 0, 4, 1, 4, 3, 1, 1, 1, 2, 1, 4, 0, 3, 0, 3, 0, 3, 0, 5, 3, 0, 0, 1, 2, 1, 2, 4, 3, 0, 1, 0, 5, 2}, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minTaps(tc.n, tc.ranges))
		})
	}
}

func minTaps(n int, ranges []int) int {
	// Keep a sorted list of (endPos,nTaps), pos and ntaps ascending.
	// When considering a tap, find the first element with an endpos >= startpos
	// in the list, this is the optimal prior tap to start from. Then check
	// whether the first non-covered (end position) of the tap would be a new
	// optimal number of taps in the list.
	type pos struct {
		end   int
		ntaps int
	}
	reach := make([]pos, 2)
	reach[0] = pos{0, 0}
	reach[1] = pos{n, math.MaxInt32}
	for i, d := range ranges {
		if d == 0 {
			continue
		}
		// Find least taps to reach start of range
		start := max(0, i-d)
		end := min(n, i+d)
		l := sort.Search(len(reach), func(i int) bool {
			return reach[i].end >= start
		})
		ntaps := reach[l].ntaps + 1
		r := sort.Search(len(reach), func(i int) bool {
			return reach[i].end >= end
		})
		// Now there's a couple of cases.
		// If insert would not provide a better result for total number of taps,
		// then there is no point in using this tap.
		if reach[r].ntaps <= ntaps {
			continue
		}

		if reach[r].end == end {
			reach[r].ntaps = ntaps
		} else {
			reach = append(reach, pos{})
			copy(reach[r+1:], reach[r:])
			reach[r] = pos{end, ntaps}
		}

		// The insert may have created a better alternative for prior positions.
		// Do some cleaning here.
		for i := r - 1; i >= 0 && reach[i].ntaps >= ntaps; i-- {
			// cut out useless element
			copy(reach[i:], reach[i+1:])
			reach = reach[:len(reach)-1]
		}
	}
	last := reach[len(reach)-1]
	if last.ntaps >= math.MaxInt32 {
		return -1
	}
	return last.ntaps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

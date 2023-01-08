package p2528maximizetheminimumpoweredcity

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPower(t *testing.T) {
	for i, tc := range []struct {
		stations []int
		r        int
		k        int
		want     int64
	}{
		{[]int{1, 2, 4, 5, 0}, 1, 2, 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxPower(tc.stations, tc.r, tc.k))
		})
	}
}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	change := make([]int, n+1)
	for i, s := range stations {
		start := max(0, i-r)
		end := min(n, i+r+1)
		change[start] += s
		change[end] -= s
	}
	state := make([]int, n)
	var curr int
	for i := range state {
		curr += change[i]
		state[i] = curr
	}

	lo, hi := 0, math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(change, state, mid, r, k) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int64(lo - 1)
}

func check(change, initialState []int, want, r, k int) bool {
	n := len(initialState)
	for i := range change {
		change[i] = 0
	}
	var extra int
	for i := range initialState {
		extra += change[i]
		curr := initialState[i] + extra
		if want > curr {
			delta := want - curr
			k -= delta
			if k < 0 {
				return false
			}
			extra += delta
			end := min(n, i+2*r+1)
			change[end] -= delta
		}
	}
	return true
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

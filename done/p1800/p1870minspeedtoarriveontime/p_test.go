package p1870minspeedtoarriveontime

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSpeedOnTime(t *testing.T) {
	for _, tc := range []struct {
		dist []int
		hour float64
		want int
	}{
		{[]int{5, 3, 4, 6, 2, 2, 7}, 10.92, 4},
		{[]int{1, 1, 100000}, 2.01, 10000000},
		{[]int{1, 3, 2}, 6, 1},
		{[]int{1, 3, 2}, 2.7, 3},
		// {[]int{1, 3, 2}, 1.9, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dist), func(t *testing.T) {
			require.Equal(t, tc.want, minSpeedOnTime(tc.dist, tc.hour))
		})
	}
}

func minSpeedOnTime(dist []int, hour float64) int {
	// Take first train on hour zero
	// Depending on speed, we may need to wait for the second train
	// The fastest possible time is to jump on each train within an hour
	n := len(dist)
	if hour < float64(n-1) {
		return -1
	}
	// Otherwise, it is always possible
	// Let's try binary search O(nlogn)
	res := sort.Search(10000000, func(i int) bool {
		if i == 0 {
			return false
		}
		res := calcTime(dist, n, i)
		return res <= hour
	})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcTime(dist []int, n, speed int) float64 {
	var res int
	for i := 0; i < n-1; i++ {
		res += dist[i] / speed
		if dist[i]%speed != 0 {
			res++
		}
	}
	final := float64(res) + float64(dist[n-1])/float64(speed)
	return final
}

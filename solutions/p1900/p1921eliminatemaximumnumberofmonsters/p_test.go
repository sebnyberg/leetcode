package p1921eliminatemaximumnumberofmonsters

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_eliminateMaximum(t *testing.T) {
	for _, tc := range []struct {
		dist  []int
		speed []int
		want  int
	}{
		{[]int{1, 3, 4}, []int{1, 1, 1}, 3},
		{[]int{3, 5, 7, 4, 5}, []int{2, 3, 6, 3, 2}, 2},
		{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}, 1},
		{[]int{3}, []int{5}, 1},
		{[]int{1, 3, 4}, []int{1, 1, 1}, 3},
		{[]int{3, 2, 4}, []int{5, 3, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dist), func(t *testing.T) {
			require.Equal(t, tc.want, eliminateMaximum(tc.dist, tc.speed))
		})
	}
}

func eliminateMaximum(dist []int, speed []int) int {
	// Divide dist by speed
	div := make([]int, len(dist))
	for i := range div {
		div[i] = dist[i] / speed[i]
		if dist[i]%speed[i] != 0 {
			div[i]++
		}
	}
	// Count number of monsters per speed in div
	count := make(map[int]int)
	var maxVal int
	minVal := math.MaxInt32
	for _, val := range div {
		minVal = min(minVal, val)
		maxVal = max(maxVal, val)
		count[val]++
	}
	// Remove first enemy
	count[minVal]--
	neliminated := 1
	if count[0] > 0 {
		return 1
	}
	ncaneliminate := 0
	for i := 1; i <= maxVal; i++ {
		val := count[i]
		if val > ncaneliminate {
			return neliminated + ncaneliminate
		} else {
			neliminated += val
			ncaneliminate -= val
		}
		ncaneliminate++
	}
	return neliminated
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

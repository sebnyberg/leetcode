package p1578minimumtimetomakeropecolorful

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		colors     string
		neededTime []int
		want       int
	}{
		{"abaac", []int{1, 2, 3, 4, 5}, 3},
		{"abc", []int{1, 2, 3}, 0},
		{"aabaa", []int{1, 2, 3, 4, 1}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.colors, tc.neededTime))
		})
	}
}

func minCost(colors string, neededTime []int) int {
	// From what I can gather, this has a greedy solution.
	//
	// I.e., when there are two consecutive balloons, at least one of them will
	// have to be removed. For a series of balloons, the largest neededTime[i]
	// balloon should be kept to minimize total time taken.
	//
	var res, sum, maxVal int
	for i := 0; i < len(colors); i++ {
		sum += neededTime[i]
		maxVal = max(maxVal, neededTime[i])
		if i == len(colors)-1 || colors[i] != colors[i+1] {
			res += sum - maxVal
			sum = 0
			maxVal = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

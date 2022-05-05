package p2048nextgreaternumericallybalancednumber

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextBeautifulNumber(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 22},
		{1000, 1333},
		{3000, 3133},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nextBeautifulNumber(tc.n))
		})
	}
}
func nextBeautifulNumber(n int) int {
	// There are 6! = 720 different ways of picking numbers.
	// Then 720 different ways of combining them (given n < 10^7)
	f := minAboveFinder{
		minAbove:    10000000,
		minAboveLen: 7,
	}
	valPicked := make([]bool, 7)
	f.dfs(1, 6, n+1, valPicked)
	return f.minAbove
}

type minAboveFinder struct {
	minAbove    int
	minAboveLen int
}

func (f *minAboveFinder) dfs(cur, max, want int, valPicked []bool) {
	if cur > max {
		// Traverse mutations of picks and check which one was optimal
		totalCount := 0
		alternatives := make([]int, 0, totalCount)
		for n, p := range valPicked {
			if p {
				totalCount += n
				for i := 0; i < n; i++ {
					alternatives = append(alternatives, n)
				}
			}
		}
		// Skip infeasible candidates
		if totalCount > f.minAboveLen {
			return
		}
		picked := make([]bool, totalCount)
		f.minAbove = min(f.minAbove, dfs2(0, totalCount, 0, want, alternatives, picked))
		return
	}
	valPicked[cur] = true // pick
	f.dfs(cur+1, max, want, valPicked)
	valPicked[cur] = false // don't pick
	f.dfs(cur+1, max, want, valPicked)
}

func dfs2(curIdx, totalCount, currentVal, want int, alternatives []int, picked []bool) int {
	if curIdx == totalCount {
		if currentVal >= want {
			return currentVal
		}
		return math.MaxInt32
	}
	currentVal *= 10
	minAbove := math.MaxInt32
	for i := range picked {
		if !picked[i] {
			// Try picking
			picked[i] = true
			currentVal += alternatives[i]
			minAbove = min(minAbove, dfs2(curIdx+1, totalCount, currentVal, want, alternatives, picked))
			// Unset
			picked[i] = false
			currentVal -= alternatives[i]
		}
	}
	return minAbove
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p1388pizzawith3nslices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSizeSlices(t *testing.T) {
	for i, tc := range []struct {
		slices []int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 10},
		{[]int{8, 9, 8, 6, 1, 1}, 16},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxSizeSlices(tc.slices))
		})
	}
}

func maxSizeSlices(slices []int) int {
	// The difficult part with this problem is somehow figuring out that it is
	// possible to pick any set of n/3 slices so long as two slices are not
	// adjacent. I don't know how people figured this out - I ran it on paper
	// but could not prove why it always works.
	//
	// To avoid picking the first AND last element, we can run the algorithm
	// twice, once for the range [:n-1] and once for [1:]
	n := len(slices)
	m := n / 3
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// dp[j][k] = max total size of slices given j'th slice considered and k
	// consumed slices. The best result is given by
	// dp[len(slices)-1][len(slices)/3]
	doSlice := func(slices []int) int {
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = 0
			}
		}
		dp[0][1] = slices[0]
		for i := 1; i < len(slices); i++ {
			dp[i][1] = max(dp[i-1][1], slices[i])
		}
		for k := 2; k <= m; k++ {
			for j := (k - 1) * 2; j < len(slices); j++ {
				dp[j][k] = max(
					dp[j-1][k],
					dp[j-2][k-1]+slices[j],
				)
			}
		}
		return dp[len(slices)-1][m]
	}
	a := doSlice(slices[1:])
	b := doSlice(slices[:n-1])
	res := max(a, b)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

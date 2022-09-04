package p2400numberofwaystoreachapositionafterexactlyksteps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWays(t *testing.T) {
	for _, tc := range []struct {
		startPos int
		endPos   int
		k        int
		want     int
	}{
		{1000, 1000, 1000, 1},
		{1, 2, 3, 3},
		{2, 5, 10, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startPos), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWays(tc.startPos, tc.endPos, tc.k))
		})
	}
}

func numberOfWays(startPos int, endPos int, k int) int {
	// Bottom-up DP with clever indexing
	const offset = 500
	var curr [2002]int
	var next [2002]int
	const mod = 1e9 + 7
	startPos += offset
	endPos += offset
	curr[startPos] = 1
	for i := 1; i <= k; i++ {
		l := max(startPos-i, endPos-(k-i))
		r := min(startPos+i, endPos+(k-i))
		for j := l; j <= r; j++ {
			next[j] = (curr[j-1] + curr[j+1]) % mod
		}
		curr, next = next, curr
	}
	return curr[endPos]
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

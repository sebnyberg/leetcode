package p0062uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePaths(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, uniquePaths(tc.m, tc.n))
		})
	}
}

func uniquePaths(m int, n int) int {
	// The way to construct the solution below is to go from a memoization of size
	// NxN to 2xN to N. Without those steps, it will be very hard to understand
	// how the solution was conceived.
	var nways [101]int
	nways[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			nways[j] = nways[j-1] + nways[j]
		}
	}
	return nways[n-1]
}

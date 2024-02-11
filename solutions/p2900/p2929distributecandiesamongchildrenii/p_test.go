package p2929distributecandiesamongchildrenii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distributeCandies(t *testing.T) {
	for i, tc := range []struct {
		n     int
		limit int
		want  int64
	}{
		{5, 2, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, distributeCandies(tc.n, tc.limit))
		})
	}
}

func distributeCandies(n int, limit int) int64 {
	f := func(vals []int) []int {
		presum := make([]int, len(vals)+1)
		for i := range vals {
			presum[i+1] = presum[i] + vals[i]
		}
		return presum
	}
	curr := make([]int, n+1)
	curr[n] = 1

	for i := 0; i < 3; i++ {
		next := make([]int, n+1)
		presum := f(curr)
		for k := 0; k <= n; k++ {
			upper := min(k+limit, n)
			next[k] = presum[upper+1] - presum[k]
		}
		curr = next
	}
	return int64(curr[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

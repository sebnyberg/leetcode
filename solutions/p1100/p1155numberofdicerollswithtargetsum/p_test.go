package p1155numberofdicerollswithtargetsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numRollsToTarget(t *testing.T) {
	for i, tc := range []struct {
		n      int
		k      int
		target int
		want   int
	}{
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{30, 30, 500, 222616187},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numRollsToTarget(tc.n, tc.k, tc.target))
		})
	}
}

func numRollsToTarget(n int, k int, target int) int {
	// Consider n=1, k=6. The distribution is:
	//
	//   [1,1,1,1,1,1]
	//
	// With n=2, the second round would yield:
	//
	// +   [1,1,1,1,1,1]
	// +     [1,1,1,1,1,1]
	// +       [...]
	// = [0,1,2,3,4,5,6,5,4,3,2,1]
	//
	// The number of ways to reach a certain result is given by the sum of ways
	// to reach the previous k results from the prior round.
	//
	curr := make([]int, target+1)
	currsum := make([]int, target+2)
	next := make([]int, target+1)
	nextsum := make([]int, target+2)
	curr[0] = 1
	const mod = 1e9 + 7
	presum := func(a, b []int) {
		for i := 0; i < len(b)-1; i++ {
			b[i+1] = (b[i] + a[i]) % mod
		}
	}
	presum(curr, currsum)
	for x := 1; x <= n; x++ {
		for j := 0; j < x; j++ {
			next[j] = 0
		}
		for i := x; i <= target; i++ {
			next[i] = (currsum[i] - currsum[max(0, i-k)] + mod) % mod
		}
		presum(next, nextsum)
		curr, next = next, curr
		currsum, nextsum = nextsum, currsum
	}
	return curr[target]
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

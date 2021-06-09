package p0403frogjump

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canCross(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 2147483647}, false},
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, canCross(tc.stones))
		})
	}
}

func canCross(stones []int) bool {
	// There are a maximum of 2000 stones total
	// For each stone arrival, there are 3 possible new jumps
	// It is not feasible to consider 3^2000 different jumps (exaggeration)
	// It would be OK for a O(n^2) algo which is a pairwise comparison of all
	// pairs in the array.
	// Let's try brute-force - did not work
	// Let's try DFS
	n := len(stones)
	// failed[i][k] = whether all jumps from i failed given that the previous
	// jump was of height k
	failed := make([][]bool, n)
	for i := range failed {
		failed[i] = make([]bool, n*2+1)
	}
	return dfs(failed, stones, n, 0, 0)
}

func dfs(failed [][]bool, stones []int, n, idx, k int) bool {
	// Base case
	if idx == n-1 {
		return true
	}
	if failed[idx][k] {
		return false
	}
	for j := idx + 1; j < n; j++ {
		gap := stones[j] - stones[idx]
		if gap >= k-1 && gap <= k+1 {
			if dfs(failed, stones, n, j, gap) {
				return true
			}
		}
	}
	failed[idx][k] = true
	return false
}

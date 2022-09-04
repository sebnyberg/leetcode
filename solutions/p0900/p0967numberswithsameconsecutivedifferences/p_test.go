package p0967numberswithsameconsecutivedifferences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numsSameConsecDiff(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want []int
	}{
		{3, 1, []int{101, 121, 123, 210, 212, 232, 234, 321, 323, 343, 345, 432, 434, 454, 456, 543, 545, 565, 567, 654, 656, 676, 678, 765, 767, 787, 789, 876, 878, 898, 987, 989}},
		{3, 7, []int{181, 292, 707, 818, 929}},
		{2, 0, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, numsSameConsecDiff(tc.n, tc.k))
		})
	}
}

func numsSameConsecDiff(n int, k int) []int {
	var res []int
	ok := func(x int) bool {
		return x >= 0 && x <= 9
	}

	// Enumerate all valid values with DFS
	var dfs func(x, i int)
	dfs = func(x, i int) {
		if i == n {
			res = append(res, x)
			return
		}
		// Add and subtract k
		if ok(x%10 + k) {
			dfs(x*10+x%10+k, i+1)
		}
		if k != 0 && ok(x%10-k) {
			dfs(x*10+x%10-k, i+1)
		}
	}
	for x := 1; x <= 9; x++ {
		dfs(x, 1)
	}
	return res
}

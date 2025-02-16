package p1718constructthelexicographicallylargestvalidsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constructDistancedSequence(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []int
	}{
		{5, []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
		{3, []int{3, 1, 2, 3, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, constructDistancedSequence(tc.n))
		})
	}
}

func constructDistancedSequence(n int) []int {
	res := make([]int, 2*n-1)
	dfs(0, n, 0, res)
	return res
}

func dfs(i, n, seen int, res []int) bool {
	if i == len(res) {
		return true
	}
	if res[i] != 0 {
		return dfs(i+1, n, seen, res)
	}
	// pick numbers from high to low
	for x := n; x >= 1; x-- {
		if seen&(1<<x) != 0 {
			continue
		}
		// 1 is a special case
		if x == 1 {
			res[i] = 1
			if dfs(i+1, n, seen|2, res) {
				return true
			}
			res[i] = 0
			continue
		}
		// number hasn't been picked
		if i+x >= len(res) {
			// this number will never be pickable
			return false
		}
		if res[i+x] != 0 {
			continue
		}
		// pick the number
		res[i] = x
		res[i+x] = x
		if dfs(i+1, n, seen|(1<<x), res) {
			return true
		}
		res[i] = 0
		res[i+x] = 0
	}
	return false
}

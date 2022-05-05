package p0077combinations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combine(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want [][]int
	}{
		// {4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.n, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, combine(tc.n, tc.k))
		})
	}
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	findCombinations(1, n, make([]int, 0, k), k, &res)
	return res
}

func findCombinations(min, n int, prefix []int, k int, res *[][]int) {
	if len(prefix) == k {
		*res = append(*res, prefix)
		return
	}

	maxPossibleNum := n - (k - len(prefix) - 1)
	for i := min; i <= maxPossibleNum; i++ {
		if i == maxPossibleNum {
			prefix = append(prefix, i)
			findCombinations(i+1, n, prefix, k, res)
			break
		}
		newPrefix := make([]int, len(prefix))
		copy(newPrefix, prefix)
		newPrefix = append(newPrefix, i)
		findCombinations(i+1, n, newPrefix, k, res)
	}
}

package p0039combinationsum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum(t *testing.T) {
	for _, tc := range []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.candidates, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, combinationSum(tc.candidates, tc.target))
		})
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return sortedCombinationSum(candidates, target)
}

func sortedCombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	for i, n := range candidates {
		switch {
		case n < target:
			for _, nn := range combinationSum(candidates[i:], target-n) {
				sub := make([]int, len(nn)+1)
				copy(sub[1:], nn)
				sub[0] = n
				res = append(res, sub)
			}
			continue
		case n == target:
			res = append(res, []int{n})
		}
		break
	}
	return res
}

package p0046permutations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permuteUnique(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, permuteUnique(tc.nums))
		})
	}
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	prefix := make([]int, len(nums))
	collectPerms(nums, prefix, 0, 0, len(nums), &res)
	return res
}

func collectPerms(nums, prefix []int, bm, i, n int, result *[][]int) {
	if i == n {
		cpy := make([]int, n)
		copy(cpy, prefix)
		*result = append(*result, cpy)
		return
	}
	seen := 0
	for j, num := range nums {
		nonNeg := num + 10
		if bm&(1<<j) > 0 || seen&(1<<nonNeg) > 0 {
			continue
		}
		seen |= (1 << nonNeg)
		prefix[i] = num
		collectPerms(nums, prefix, bm|(1<<j), i+1, n, result)
	}
}

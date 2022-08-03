package p0646maximumlengthofpairchain

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLongestChain(t *testing.T) {
	for _, tc := range []struct {
		pairs [][]int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[1,2],[2,3],[3,4]]"),
			2,
		},
		{
			leetcode.ParseMatrix("[[1,2],[7,8],[4,5]]"),
			3,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pairs), func(t *testing.T) {
			require.Equal(t, tc.want, findLongestChain(tc.pairs))
		})
	}
}

func findLongestChain(pairs [][]int) int {
	// Greedy will work here. Sort by end, ascending, then iterate.
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 1
	cur := pairs[0]
	for _, p := range pairs[1:] {
		if p[0] <= cur[1] {
			continue
		}
		count++
		cur = p
	}
	return count
}

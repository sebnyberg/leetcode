package p0851loudandrich

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_loudAndRich(t *testing.T) {
	for _, tc := range []struct {
		richer [][]int
		quiet  []int
		want   []int
	}{
		{
			leetcode.ParseMatrix("[[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]]"),
			[]int{3, 2, 5, 4, 6, 1, 7, 0},
			[]int{5, 5, 2, 5, 4, 5, 6, 7},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.richer), func(t *testing.T) {
			require.Equal(t, tc.want, loudAndRich(tc.richer, tc.quiet))
		})
	}
}

func loudAndRich(richer [][]int, quiet []int) []int {
	// Point rich people toward poor people
	// Then perform topological search, updating the minimum quietness of each
	// group for a given node
	n := len(quiet)
	adj := make([][]int, n)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	indeg := make([]int, n)
	for _, r := range richer {
		// a has more money than b
		a, b := r[0], r[1]
		indeg[b]++
		adj[a] = append(adj[a], b)
	}
	curr := []int{}
	for i := range indeg {
		if indeg[i] == 0 {
			curr = append(curr, i)
		}
	}
	next := []int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, nei := range adj[x] {
				indeg[nei]--
				if quiet[ans[x]] < quiet[ans[nei]] {
					ans[nei] = ans[x]
				}
				if indeg[nei] == 0 {
					next = append(next, nei)
				}
			}
		}
		curr, next = next, curr
	}
	return ans
}

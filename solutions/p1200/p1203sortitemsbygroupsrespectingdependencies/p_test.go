package p1203sortitemsbygroupsrespectingdependencies

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_sortItems(t *testing.T) {
	for i, tc := range []struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
		want        []int
	}{
		{
			8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1},
			leetcode.ParseMatrix("[[],[6],[5],[6],[3,6],[],[7],[]]"),
			[]int{6, 3, 4, 1, 5, 2, 0, 7},
		},
		{
			8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1},
			leetcode.ParseMatrix("[[],[6],[5],[6],[3],[],[4],[]]"),
			[]int{},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			groupItems := make(map[int][]int)
			for i := range tc.group {
				gi := tc.group[i]
				groupItems[gi] = append(groupItems[gi], i)
			}
			res := sortItems(tc.n, tc.m, tc.group, tc.beforeItems)

			// validate that group ordering is correct
			var want int
			var currGroup int
			for i := range res {
				gi := tc.group[res[i]]
				if gi == -1 {
					require.Equal(t, want, 0,
						"idx:%v, wanted group item but got ungrouped number", i)
					continue
				}
				if want == 0 {
					want = len(groupItems[gi]) - 1
					currGroup = gi
					continue
				}
				require.Equal(t, currGroup, gi,
					"idx:%v, current group did not match expected", i)
				want--
			}

			// validate that numbers come in their expected order
			seen := make([]bool, tc.n)
			for i := range res {
				for _, x := range tc.beforeItems[res[i]] {
					require.True(t, seen[x],
						"idx:%v missing requirement %v", i, x,
					)
				}
				seen[res[i]] = true
			}
		})
	}
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// One of the easiest but most pain-staking problems on Leetcode..

	// Assign single items to groups
	for i := range group {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	// Create group-based adj-list
	indeg := make([]int, m)
	adj := make([][]int, m)
	for i, xs := range beforeItems {
		for _, x := range xs {
			gx := group[x]
			gi := group[i]
			if gx == gi {
				continue
			}
			adj[gx] = append(adj[gx], gi) // placing gx leads to gi
		}
	}

	// De-duplicade adj and calculate indegree
	for i := range adj {
		if len(adj[i]) == 0 {
			continue
		}
		sort.Ints(adj[i])
		var j int
		for k := range adj[i] {
			if adj[i][j] == adj[i][k] {
				continue
			}
			j++
			adj[i][j] = adj[i][k]
		}
		adj[i] = adj[i][:j+1]
	}
	for i := range adj {
		for _, x := range adj[i] {
			indeg[x]++
		}
	}

	k := 0 // seen groups
	// Topo-sort
	curr := []int{}
	next := []int{}
	var j int
	groupIdx := make([]int, m)
	for g, deg := range indeg {
		if deg == 0 {
			groupIdx[g] = j
			j++
			curr = append(curr, g)
			k++
		}
	}

	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, y := range adj[x] {
				indeg[y]--
				if indeg[y] == 0 {
					groupIdx[y] = j
					j++
					next = append(next, y)
					k++
				}
			}
		}

		curr, next = next, curr
	}
	if k != m {
		return []int{}
	}

	// At this point it is known that groups are sortable.
	// Now it is time for each group to be internally sorted, once again, using
	// topo-sort but for all numbers at once. A topo-sorted number is added to
	// its respective group and accounted for.
	adj = append(adj[:0], make([][]int, n)...)
	indeg = append(indeg[:0], make([]int, n)...)
	for i, xs := range beforeItems {
		for _, x := range xs {
			adj[x] = append(adj[x], i)
			indeg[i]++
		}
	}
	curr = curr[:0]
	k = 0
	res := make([][]int, m)
	for x, deg := range indeg {
		if deg == 0 {
			curr = append(curr, x)
			gi := groupIdx[group[x]]
			res[gi] = append(res[gi], x)
			k++
		}
	}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, y := range adj[x] {
				indeg[y]--
				if indeg[y] == 0 {
					k++
					gi := groupIdx[group[y]]
					res[gi] = append(res[gi], y)
					next = append(next, y)
				}
			}
		}

		curr, next = next, curr
	}
	if k != n {
		return []int{}
	}

	var flattened []int
	for _, r := range res {
		flattened = append(flattened, r...)
	}
	return flattened
}

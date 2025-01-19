package p3203findminimumdiameteraftermergingtwotrees

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumDiameterAfterMerge(t *testing.T) {
	for _, tc := range []struct {
		edges1 [][]int
		edges2 [][]int
		want   int
	}{
		{leetcode.ParseMatrix("[[0,1],[0,2],[0,3]]"), leetcode.ParseMatrix("[[0,1]]"), 3},
		{leetcode.ParseMatrix("[[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]]"), leetcode.ParseMatrix("[[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]]"), 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges1), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDiameterAfterMerge(tc.edges1, tc.edges2))
		})
	}
}

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	// This is rather simple. We can process nodes by indegree (1 == edge) to
	// calculate the distance to a middle node in a tree. Then just add the
	// distances together.
	m := len(edges1) + 1
	n := len(edges2) + 1
	indeg := make([]int, 0, max(m, n))
	adj := make([][]int, 0, max(m, n))
	seen := make([]bool, 0, max(m, n))
	var curr []int
	var next []int

	findMiddleNodes := func(edges [][]int) (radius int, nmiddle int) {
		if len(edges) == 0 {
			return 0, 1
		}
		adj = adj[:len(edges)+1]
		indeg = indeg[:len(edges)+1]
		seen = seen[:len(edges)+1]
		for i := range seen {
			// reset
			seen[i] = false
			adj[i] = adj[i][:0]
			indeg[i] = 0
		}
		for _, e := range edges {
			a := e[0]
			b := e[1]
			indeg[a]++
			indeg[b]++
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
		curr = curr[:0]
		next = next[:0]
		for i, d := range indeg {
			if d == 1 {
				curr = append(curr, i)
			}
		}
		radius = -1
		for ; len(curr) > 0; radius++ {
			next = next[:0]
			for _, x := range curr {
				for _, u := range adj[x] {
					indeg[u]--
					if indeg[u] == 1 {
						next = append(next, u)
					}
				}
			}
			curr, next = next, curr
		}

		middleNodes := next
		return radius, len(middleNodes)
	}
	aRadius, aNumMiddles := findMiddleNodes(edges1)
	bRadius, bNumMiddles := findMiddleNodes(edges2)

	// The tricky part is that if the choice of middle node is arbitrary, then the
	// radius is off-by-one compared to when the choice is is exact. That is, if
	// the radius is 3 and there are two middle nodes, then the actual diameter is
	// 3 + 3 + 1(edge between middle nodes).
	aMaxRadius := aRadius
	bMaxRadius := bRadius
	if aNumMiddles > 1 {
		aMaxRadius++
	}
	if bNumMiddles > 1 {
		bMaxRadius++
	}
	aDiam := aRadius + aMaxRadius
	bDiam := bRadius + bMaxRadius
	mergedDiam := aMaxRadius + bMaxRadius + 1

	return max(mergedDiam, max(aDiam, bDiam))
}

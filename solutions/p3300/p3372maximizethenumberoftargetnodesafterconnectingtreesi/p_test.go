package p3372maximizethenumberoftargetnodesafterconnectingtreesi

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxTargetNodes(t *testing.T) {
	for _, tc := range []struct {
		edges1 [][]int
		edges2 [][]int
		k      int
		want   []int
	}{
		{
			leetcode.ParseMatrix("[[0,1],[0,2],[2,3],[2,4]]"),
			leetcode.ParseMatrix("[[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]"),
			2,
			[]int{9, 7, 9, 8, 8},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges1), func(t *testing.T) {
			require.Equal(t, tc.want, maxTargetNodes(tc.edges1, tc.edges2, tc.k))
		})
	}
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	var n1 int
	var n2 int
	for i := range edges1 {
		n1 = max(n1, edges1[i][0], edges1[i][1])
	}
	for i := range edges2 {
		n2 = max(n2, edges2[i][0], edges2[i][1])
	}
	n1++
	n2++
	if k == 0 {
		res := make([]int, n1)
		for i := range res {
			res[i] = 1
		}
		return res
	}

	var curr []int
	var next []int
	f := func(edges [][]int, n, k int) ([]int, int) {
		// Capture edges as adjacency list
		adj := make([][]int, n)
		seen := make([]bool, n)
		var maxVisitable int
		for _, e := range edges {
			a := e[0]
			b := e[1]
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
		visitable := make([]int, n)
		for i := range adj {
			curr = append(curr[:0], i)
			seen = append(seen[:0], make([]bool, n)...)
			seen[i] = true
			res := 1
			for kk := k; len(curr) > 0 && kk > 0; kk-- {
				next = next[:0]
				for _, a := range curr {
					for _, b := range adj[a] {
						if seen[b] {
							continue
						}
						seen[b] = true
						next = append(next, b)
					}
				}
				res += len(next)
				curr, next = next, curr
			}
			maxVisitable = max(maxVisitable, res)
			visitable[i] = res
		}
		return visitable, maxVisitable
	}
	a, _ := f(edges1, n1, k)
	_, maxVisitable := f(edges2, n2, k-1)
	res := make([]int, n1)
	for i := range n1 {
		res[i] = a[i] + maxVisitable
	}
	return res
}

package p2872maximumnumberofkdivisiblecomponents

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxKDivisibleComponents(t *testing.T) {
	for i, tc := range []struct {
		n      int
		edges  [][]int
		values []int
		k      int
		want   int
	}{
		{
			5,
			leetcode.ParseMatrix("[[0,2],[1,2],[1,3],[2,4]]"),
			[]int{1, 8, 1, 4, 4},
			6,
			2,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxKDivisibleComponents(tc.n, tc.edges, tc.values, tc.k))
		})
	}
}

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	// Whenever there is a viable cut, then we should make it.
	//
	// I'm thinking that if we do BFS from nodes with indegree 1, then we
	// attribute the value of each node to its parent. Whenever we visit a node
	// that is evenly divisible by k, we increment the total count.
	if len(values) == 1 {
		if values[0]%k == 0 {
			return 1
		}
		return 0
	}

	// Build adjacency list
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		indeg[u]++
		indeg[v]++
	}

	curr := []int{}
	next := []int{}
	for i := range values {
		if indeg[i] == 1 {
			indeg[i]--
			curr = append(curr, i)
		}
	}
	var res int
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			if values[x]%k == 0 {
				res++
			}
			for _, v := range adj[x] {
				indeg[v]--
				values[v] += values[x]
				if indeg[v] == 1 {
					next = append(next, v)
				}
			}
		}
		curr, next = next, curr
	}
	return res
}

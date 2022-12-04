package p2493dividenotesintothemaximumnumberofgroups

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_magnificentSets(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		want  int
	}{
		{
			6,
			leetcode.ParseMatrix("[[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]"),
			4,
		},
		{
			3,
			leetcode.ParseMatrix("[[1,2],[2,3],[3,1]]"),
			-1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, magnificentSets(tc.n, tc.edges))
		})
	}
}

func magnificentSets(n int, edges [][]int) int {
	// The goal is to perform a Bipartite matching of the nodes in the graph(s).
	//
	// One way of doing this is to do coloring: for each node in the graph,
	// ensure that all its neighbours are either uncolored, or has a different
	// color than the node itself. For example, if a node has been assigned the
	// color "blue", then every neighbour must either be uncolored, or have the
	// color "red".
	//
	// We can make an adjustment to this algorithm and use a monotonic clock
	// instead. It the current time is, say, 4, then any valid neighbour must've
	// been seen at an odd time or not seen yet.
	//
	// The highest time assigned to a node during this traversal is the most
	// distant node group according to that node. For any graph with a valid
	// matching, these must exist at least two nodes for which the distance to
	// the most distant node group is the number of node groups in total.
	//
	// To keep track of which nodes belong to which graph, we can use a DSU.
	//

	// DSU
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		ra := find(parent[a])
		parent[a] = ra
		return ra
	}
	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra != rb {
			parent[rb] = ra
		}
	}

	// Form adjacency list, group nodes in DSU
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
		union(e[0], e[1])
	}

	curr := []int{}
	next := []int{}
	times := make([]uint16, n+1)
	longestDistanceFromNode := func(i int) int {
		// "Color" the graph using a monotic clock
		for i := range times {
			times[i] = math.MaxUint16
		}
		curr := curr[:0]
		curr = append(curr, i)
		var t uint16
		times[i] = 0
		for len(curr) > 0 {
			t++
			next = next[:0]
			for _, x := range curr {
				for _, y := range adj[x] {
					if times[y] == math.MaxUint16 {
						times[y] = t
						next = append(next, y)
						continue
					}
					// Any previously visited node must have the expected
					// "color", i.e. odd or even timestamp
					if times[y]&1 != t&1 {
						return -1
					}
				}
			}
			curr, next = next, curr
		}
		return int(t)
	}

	ngroups := make([]int, n+1)
	var res int
	for i := 1; i <= n; i++ {
		dist := longestDistanceFromNode(i)
		if dist == -1 {
			return -1
		}
		// Recall: for at least two nodes, the number of groups is the same
		// thing as the distance
		root := find(i) // Use DSU to bundle graph nodes together
		ngroups[root] = max(ngroups[root], dist)
	}
	for _, c := range ngroups {
		res += c
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

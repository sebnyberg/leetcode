package p2867countvalidpathsinatree

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countPaths(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		want  int64
	}{
		{5, leetcode.ParseMatrix("[[1,2],[1,3],[2,4],[2,5]]"), 4},
		{6, leetcode.ParseMatrix("[[1,2],[1,3],[2,4],[3,5],[3,6]]"), 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countPaths(tc.n, tc.edges))
		})
	}
}

func countPaths(n int, edges [][]int) int64 {
	// Imagine cutting the tree at any prime node. The result will be a bunch of
	// subtrees containing no primes.
	//
	// From the point of view of any given prime node, the total number of paths
	// is given by combining possible paths from each non-prime subtree. It
	// turns out that the total number of such paths for one subtree is equal to
	// the number of nodes in the subtree (try this on paper).
	//
	// This means that we could iterate over all non-prime subtrees one-by-one,
	// adding the number of nodes in the subtree, and the number of nodes in all
	// previously considered subtrees.
	//
	// To find the size of each subtree, we can join together all non-prime
	// pairs into the same group using a modified union-find with group sizes.
	//
	size := make([]int, n+1)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		parent[a] = find(parent[a])
		return parent[a]
	}

	union := func(a, b int) {
		a, b = find(a), find(b)
		if a != b {
			size[b] += size[a]
			parent[a] = b
		}
	}

	// Find primes from 1 to n
	seen := make([]bool, n+1)
	isprime := make(map[int]bool) // cba to do struct{}

	// Sieve of eratosthenes
	for i := 2; i <= n; i++ {
		if seen[i] {
			continue
		}
		isprime[i] = true
		seen[i] = true
		for k := i; k <= n; k += i {
			seen[k] = true
		}
	}

	// Build adjacency list and union-find non-prime pairs
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		if isprime[u] || isprime[v] {
			continue
		}
		union(u, v) // both non-primes
	}

	var res int
	for prime := range isprime {
		// For each non-prime neighbour
		var sz int
		for _, v := range adj[prime] {
			// Add the number of nodes in the subtree + multiply by prior nodes
			if isprime[v] {
				continue
			}
			treesz := size[find(v)]
			res += (sz + 1) * treesz
			sz += treesz
		}
	}

	return int64(res)
}

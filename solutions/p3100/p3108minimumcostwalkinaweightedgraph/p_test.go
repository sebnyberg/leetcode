package p3108minimumcostwalkinaweightedgraph

import "math"

func minimumCost(n int, edges [][]int, query [][]int) []int {
	parent := make([]int, n+1)
	val := make([]int, n+1)
	for i := range parent {
		parent[i] = i
		val[i] = math.MaxInt32
	}
	var find func(int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		root := find(parent[a])
		parent[a] = root
		return root
	}
	union := func(a, b, w int) {
		val[a] &= w
		val[b] &= w
		a = find(a)
		b = find(b)
		if a != b {
			parent[b] = a
			val[a] &= val[b]
		}
		val[a] &= w
	}
	for _, e := range edges {
		union(e[0], e[1], e[2])
	}
	res := make([]int, len(query))
	for i := range res {
		a := find(query[i][0])
		b := find(query[i][1])
		if a != b { // not in the same graph
			res[i] = -1
		} else { // in the same graph, the min cost is found at the root
			res[i] = val[find(a)]
		}
	}
	return res
}

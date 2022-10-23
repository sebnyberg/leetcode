package p2421numberofgoodpaths

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numberOfGoodPaths(t *testing.T) {
	for i, tc := range []struct {
		vals  []int
		edges [][]int
		want  int
	}{
		{[]int{1, 3, 2, 1, 3}, leetcode.ParseMatrix("[[0,1],[0,2],[2,3],[2,4]]"), 6},
		{[]int{1, 1, 2, 2, 3}, leetcode.ParseMatrix("[[0,1],[1,2],[2,3],[2,4]]"), 7},
		{[]int{1}, [][]int{}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfGoodPaths(tc.vals, tc.edges))
		})
	}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	type node struct {
		idx int
		val int
		nei []int
	}
	nodes := make([]node, n)
	for i := range vals {
		nodes[i].val = vals[i]
		nodes[i].idx = i
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val
	})
	nodeIdx := make([]int, n)
	for i, e := range nodes {
		nodeIdx[e.idx] = i
	}
	for _, e := range edges {
		a, b := nodeIdx[e[0]], nodeIdx[e[1]]
		nodes[a].nei = append(nodes[a].nei, b)
		nodes[b].nei = append(nodes[b].nei, a)
	}
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
	m := make(map[int]int)
	res := n
	for i := range nodes {
		if i == n-1 || nodes[i+1].val != nodes[i].val {
			for k := range m {
				delete(m, k)
			}
			for a := i; ; a-- {
				for _, b := range nodes[a].nei {
					if nodes[b].val <= nodes[a].val {
						union(a, b)
					}
				}
				if a == 0 || nodes[a-1].val != nodes[a].val {
					break
				}
			}
			for a := i; ; a-- {
				m[find(a)]++
				if a == 0 || nodes[a-1].val != nodes[a].val {
					break
				}
			}
			for _, count := range m {
				res += count * (count - 1) / 2
			}
		}
	}
	return res
}

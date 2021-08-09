package p1938maximumgeneticdifferencequery

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxGeneticDifference(t *testing.T) {
	for _, tc := range []struct {
		parents []int
		queries [][]int
		want    []int
	}{
		{[]int{-1, 0, 1, 1}, [][]int{{0, 2}, {3, 2}, {2, 5}}, []int{2, 3, 7}},
		{[]int{3, 7, -1, 2, 0, 7, 0, 2}, [][]int{{4, 6}, {1, 15}, {0, 5}}, []int{6, 14, 7}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.parents), func(t *testing.T) {
			require.Equal(t, tc.want, maxGeneticDifference(tc.parents, tc.queries))
		})
	}
}

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

func (n *TrieNode) Add(x, incr int) {
	cur := n
	for i := 17; i >= 0; i-- {
		val := (x >> i) & 1
		if cur.children[val] == nil {
			cur.children[val] = new(TrieNode)
		}
		cur.children[val].count += incr
		cur = cur.children[val]
	}
}

func (n *TrieNode) FindMaxXOR(x int) int {
	cur := n
	res := 0
	for i := 17; i >= 0; i-- {
		val := (x >> i) & 1
		if ch := cur.children[1-val]; ch != nil && ch.count > 0 {
			res |= 1 << i
			cur = ch
		} else if cur.children[val] != nil {
			cur = cur.children[val]
		} else {
			break
		}
	}
	return res
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	// Instead of iterating over queries, iterate over all possible paths in the
	// graph starting with the root node.
	n := len(parents)
	adj := make([][]int, n)
	var rootIdx int
	for i, parent := range parents {
		if parent == -1 {
			rootIdx = i
			continue
		}
		adj[parent] = append(adj[parent], i)
	}
	rootNode := &TrieNode{}
	qs := make(map[int][]*nodeQuery)
	for i, query := range queries {
		node, val := query[0], query[1]
		qs[node] = append(qs[node], &nodeQuery{
			idx:     i,
			nodeIdx: node,
			val:     val,
			result:  0,
		})
	}
	dfs(rootNode, adj, rootIdx, qs)
	res := make([]int, len(queries))
	for _, q := range qs {
		for _, qq := range q {
			res[qq.idx] = qq.result
		}
	}
	return res
}

func dfs(rootNode *TrieNode, adj [][]int, idx int, nodeQueries map[int][]*nodeQuery) {
	// Add current node to Trie
	rootNode.Add(idx, 1)

	// For each query on this node
	for _, q := range nodeQueries[idx] {
		// Find max XOR value and add it to the query
		res := rootNode.FindMaxXOR(q.val)
		q.result = res
	}

	for _, nei := range adj[idx] {
		dfs(rootNode, adj, nei, nodeQueries)
	}

	// Remove current node from trie
	rootNode.Add(idx, -1)
}

type nodeQuery struct {
	idx     int
	nodeIdx int
	val     int
	result  int
}

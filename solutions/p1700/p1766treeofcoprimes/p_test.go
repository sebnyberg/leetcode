package p1766treeofcoprimes

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_getCoprimes(t *testing.T) {
	for i, tc := range []struct {
		nums  []int
		edges [][]int
		want  []int
	}{
		{
			[]int{2, 3, 3, 2},
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3]]"),
			[]int{-1, 0, 0, 1},
		},
		{
			[]int{9, 16, 30, 23, 33, 35, 9, 47, 39, 46, 16, 38, 5, 49, 21, 44, 17, 1, 6, 37, 49, 15, 23, 46, 38, 9, 27, 3, 24, 1, 14, 17, 12, 23, 43, 38, 12, 4, 8, 17, 11, 18, 26, 22, 49, 14, 9},
			leetcode.ParseMatrix("[[17,0],[30,17],[41,30],[10,30],[13,10],[7,13],[6,7],[45,10],[2,10],[14,2],[40,14],[28,40],[29,40],[8,29],[15,29],[26,15],[23,40],[19,23],[34,19],[18,23],[42,18],[5,42],[32,5],[16,32],[35,14],[25,35],[43,25],[3,43],[36,25],[38,36],[27,38],[24,36],[31,24],[11,31],[39,24],[12,39],[20,12],[22,12],[21,39],[1,21],[33,1],[37,1],[44,37],[9,44],[46,2],[4,46]]"),
			[]int{-1, 21, 17, 43, 10, 42, 7, 13, 29, 44, 17, 31, 39, 10, 10, 29, 32, 0, 40, 23, 12, 39, 12, 40, 25, 35, 15, 38, 40, 40, 17, 24, 5, 1, 19, 14, 17, 21, 25, 24, 14, 17, 40, 25, 37, 17, 10},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getCoprimes(tc.nums, tc.edges))
		})
	}
}

type node struct {
	val      int
	idx      int
	children []*node
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// Construct the tree
	seen := make([]bool, n)
	root := &node{val: nums[0], idx: 0}
	q := []*node{root}
	seen[0] = true
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range adj[x.idx] {
			if seen[y] {
				continue
			}
			seen[y] = true
			a := &node{
				val: nums[y],
				idx: y,
			}
			x.children = append(x.children, a)
			q = append(q, a)
		}
	}

	// DFS with a path to the root
	res := make([]int, n)
	path := make([]int, n)
	latestIndex := map[int]int{}
	dfs(root, nums, path, latestIndex, 0, res)
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func dfs(cur *node, nums, path []int, latestIndex map[int]int, i int, res []int) {
	if cur.idx == 2 {
		fmt.Sprint("haha")
	}
	j := -1
	for val, pathIdx := range latestIndex {
		if gcd(val, cur.val) == 1 {
			j = max(j, pathIdx)
		}
	}
	if j != -1 {
		res[cur.idx] = path[j]
	} else {
		res[cur.idx] = -1
	}
	path[i] = cur.idx
	prevIdx := -1
	if k, exists := latestIndex[cur.val]; exists {
		prevIdx = k
	}
	latestIndex[cur.val] = i
	for _, child := range cur.children {
		dfs(child, nums, path, latestIndex, i+1, res)
	}
	if prevIdx == -1 {
		delete(latestIndex, cur.val)
	} else {
		latestIndex[cur.val] = prevIdx
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

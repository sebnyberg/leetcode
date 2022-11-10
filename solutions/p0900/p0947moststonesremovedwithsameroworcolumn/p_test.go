package p0947moststonesremovedwithsameroworcolumn

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_removeStones(t *testing.T) {
	for i, tc := range []struct {
		stones [][]int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[5,9],[9,0],[0,0],[7,0],[4,3],[8,5],[5,8],[1,1],[0,6],[7,5],[1,6],[1,9],[9,4],[2,8],[1,3],[4,2],[2,5],[4,1],[0,2],[6,5]]"),
			19,
		},
		{
			leetcode.ParseMatrix("[[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]"),
			5,
		},
		{
			leetcode.ParseMatrix("[[0,0],[0,2],[1,1],[2,0],[2,2]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[0,0]]"),
			0,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, removeStones(tc.stones))
		})
	}
}

func removeStones(stones [][]int) int {
	n := len(stones)
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(int) int
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
			size[ra] += size[rb]
		}
	}
	xs := make(map[int][]int)
	ys := make(map[int][]int)
	for i, s := range stones {
		for _, j := range xs[s[0]] {
			union(i, j)
		}
		for _, j := range ys[s[1]] {
			union(i, j)
		}
		xs[s[0]] = append(xs[s[0]], i)
		ys[s[1]] = append(ys[s[1]], i)
	}
	seen := make(map[int]bool)
	var res int
	for _, p := range parent {
		a := find(p)
		if !seen[a] {
			res += size[a] - 1
		}
		seen[a] = true
	}
	return res
}

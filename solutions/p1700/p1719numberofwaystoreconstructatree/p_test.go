package p1719numberofwaystoreconstructatree

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_checkWays(t *testing.T) {
	for i, tc := range []struct {
		pairs [][]int
		want  int
	}{
		{leetcode.ParseMatrix("[[1,2],[2,3]]"), 1},
		{leetcode.ParseMatrix("[[1,2],[2,3],[1,3]]"), 2},
		{leetcode.ParseMatrix("[[1,2],[2,3],[2,4],[1,5]]"), 0},
		{leetcode.ParseMatrix("[[3,4],[2,3],[4,5],[2,4],[2,5],[1,5],[1,4]]"), 0},
		{leetcode.ParseMatrix("[[3,5],[4,5],[2,5],[1,5],[1,4],[2,4]]"), 1},
		{leetcode.ParseMatrix("[[1,5],[1,3],[2,3],[2,4],[3,5],[3,4]]"), 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, checkWays(tc.pairs))
		})
	}
}

func checkWays(pairs [][]int) int {
	var s solver
	var maxVal int
	for _, p := range pairs {
		maxVal = max(maxVal, p[0])
		maxVal = max(maxVal, p[1])
	}
	s.nodeCount = make(map[int]int, maxVal)
	s.parent = make([]int, maxVal+1)
	return s.check(pairs)
}

type solver struct {
	nodeCount map[int]int
	parent    []int
}

func (s *solver) reset() {
	for i := range s.parent {
		s.parent[i] = i
	}
	for k := range s.nodeCount {
		delete(s.nodeCount, k)
	}
}

func (s *solver) check(pairs [][]int) int {
	s.reset()

	for k := range s.nodeCount {
		delete(s.nodeCount, k)
	}
	for _, p := range pairs {
		s.nodeCount[p[0]]++
		s.nodeCount[p[1]]++
	}

	// A root must have edges toward the all other nodes
	res := 1
	root := -1
	m := len(s.nodeCount)
	for node, cnt := range s.nodeCount {
		if cnt == m-1 {
			if root == -1 {
				root = node
			} else {
				res = 2
			}
		}
	}
	if root == -1 {
		return 0
	}

	// Find strongly connected components in the graph
	for _, p := range pairs {
		if p[0] == root || p[1] == root {
			continue
		}
		s.union(p[0], p[1])
	}

	// Sort pairs according to the pair's group
	// Root-pairs are put at the end
	sort.Slice(pairs, func(i, j int) bool {
		a := pairs[i]
		b := pairs[j]
		if a[0] == root || a[1] == root {
			return false
		}
		if b[0] == root || b[1] == root {
			return true
		}
		return s.find(a[0]) < s.find(b[0])
	})

	var groups [][][]int
	var i int
	for j, p := range pairs {
		if j > 0 &&
			(s.find(pairs[j-1][0]) != s.find(p[0]) || p[0] == root || p[1] == root) {
			groups = append(groups, pairs[i:j])
			i = j
		}
		if p[0] == root || p[1] == root {
			break
		}
	}
	for _, g := range groups {
		subres := s.check(g)
		if subres == 0 {
			return 0
		}
		res = max(res, subres)
	}
	return res
}

func (s *solver) find(a int) int {
	if s.parent[a] == a {
		return a
	}
	ra := s.find(s.parent[a])
	s.parent[a] = ra
	return ra
}

func (s *solver) union(a, b int) {
	ra := s.find(a)
	rb := s.find(b)
	s.parent[ra] = rb
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

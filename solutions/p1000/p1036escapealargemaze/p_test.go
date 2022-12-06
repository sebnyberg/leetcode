package p1036escapealargemaze

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_isEscapePossible(t *testing.T) {
	for i, tc := range []struct {
		blocked [][]int
		source  []int
		target  []int
		want    bool
	}{
		{
			leetcode.ParseMatrix("[[0,1],[1,0]]"),
			[]int{0, 0},
			[]int{0, 2},
			false,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isEscapePossible(tc.blocked, tc.source, tc.target))
		})
	}
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	// BFS around the source and target for a while to check whether either ends
	// up with no alternatives
	m := make(map[[2]int]bool)
	for _, b := range blocked {
		m[[2]int{b[0], b[1]}] = true
	}

	seen := make(map[[2]int]bool)
	dist := func(a, b [2]int) int {
		return abs(a[0]-b[0]) + abs(a[1]-b[1])
	}

	// First do BFS on the source to check whether we a) can reach the target in
	// some amount of iterations, and b) end up with no alternatives
	curr := [][2]int{{source[0], source[1]}}
	next := [][2]int{}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	seen[curr[0]] = true
	src := [2]int{source[0], source[1]}
	tar := [2]int{target[0], target[1]}
	var k int
	var maxDist int
	for ; maxDist <= 400 && len(seen) < 200*200 && len(curr) > 0; k++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				p := [2]int{ii, jj}
				if ii < 0 || jj < 0 || ii >= 1000000 || jj >= 1000000 || m[p] {
					continue
				}
				if seen[p] {
					continue
				}
				if p == tar {
					return true
				}
				maxDist = max(maxDist, dist(src, p))
				seen[p] = true
				next = append(next, p)
			}
		}
		curr, next = next, curr
	}
	if len(curr) == 0 {
		return false
	}
	curr = curr[:0]
	for k := range seen {
		delete(seen, k)
	}
	curr = append(curr, tar)
	seen[tar] = true
	maxDist = 0
	for ; maxDist <= 400 && len(seen) < 200*200 && len(curr) > 0; k++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				p := [2]int{ii, jj}
				if ii < 0 || jj < 0 || ii >= 1000000 || jj >= 1000000 || m[p] {
					continue
				}
				if seen[p] {
					continue
				}
				if p == tar {
					return true
				}
				maxDist = max(maxDist, dist(tar, p))
				seen[p] = true
				next = append(next, p)
			}
		}
		curr, next = next, curr
	}
	if len(curr) == 0 {
		return false
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

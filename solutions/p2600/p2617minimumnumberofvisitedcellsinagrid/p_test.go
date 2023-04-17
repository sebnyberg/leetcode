package p2617minimumnumberofvisitedcellsinagrid

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumVisitedCells(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]"),
			4,
		},
		{
			leetcode.ParseMatrix("[[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[2,1,0],[1,0,0]]"),
			-1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumVisitedCells(tc.grid))
		})
	}
}

func minimumVisitedCells(grid [][]int) int {
	// Whether a move is valid or not is irrespective of the prior moves, it
	// only matters what the prior position is.
	// Each row and column could be managed in isolation from each other.
	// The main "issue" with this problem is that there are potentially many
	// elements which lead to a certain "next" element. We cannot materialize
	// all edges because that is potentially O(n^2).
	//
	// If a path is taken during BFS, then cells can be permanently eliminated
	// as targets. If we use a segment tree per row, then we can quickly check
	// whether an update is needed at all.
	//
	// What about partitioned BFS? We partition by row and col, then visit each
	// row/col to determine which new positions are now plausible targets.
	//
	// Let's try it
	m := len(grid)
	n := len(grid[0])
	if m == 1 && n == 1 {
		// stupid edge-case
		return 1
	}

	// Let's do this: each row and column has its own seg-tree
	// The segtree keeps track of whether an index has been visited or not
	// During update, any unvisited node is also added to the list of next nodes
	// For each row, create its segtree
	rows := make([]*segtree, m)

	arr := make([]bool, n)
	arr[0] = true
	for i := range grid {
		if i > 0 {
			arr[0] = false
		}
		rows[i] = newSegtree(i, arr, true)
	}

	arr = append(arr[:0], make([]bool, m)...)
	arr[0] = true
	cols := make([]*segtree, n)
	for j := range grid[0] {
		if j > 0 {
			arr[0] = false
		}
		cols[j] = newSegtree(j, arr, false)
	}

	curr := [][2]int{{0, 0}}
	next := [][2]int{}

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true

	for steps := 2; len(curr) > 0; steps++ {
		next = next[:0]
		// For each row and column, find positions that are so far unvisited and
		// reachable from at least one position
		for _, x := range curr {
			i := x[0]
			j := x[1]
			right := grid[i][j] + j
			bot := grid[i][j] + i
			if right > j {
				rows[i].Update(j+1, right, seen, &next)
			}
			if bot > i {
				cols[j].Update(i+1, bot, seen, &next)
			}
		}
		if seen[m-1][n-1] {
			return steps
		}
		curr, next = next, curr
	}
	return -1
}

type segtree struct {
	visited []bool
	n       int
	m       int
	isrow   bool
	k       int
}

func newSegtree(k int, items []bool, isrow bool) *segtree {
	n := len(items)
	m := 1
	for m < n {
		m *= 2
	}
	var s segtree
	s.m = m
	s.n = n
	s.visited = make([]bool, m*2)
	for i := range items {
		s.visited[m+i] = items[i]
	}
	for i := m + len(items); i < m*2; i++ {
		s.visited[i] = true
	}
	for i := m - 1; i >= 1; i-- {
		s.visited[i] = s.visited[i*2] && s.visited[i*2+1]
	}
	s.isrow = isrow
	s.k = k
	return &s
}

func (s *segtree) Update(qlo, qhi int, seen [][]bool, next *[][2]int) {
	s.upd(1, 0, s.m-1, qlo, qhi, seen, next)
}

func (s *segtree) upd(i, loidx, hiidx, qlo, qhi int, seen [][]bool, next *[][2]int) {
	// i = current index in tree
	// [lo,hi] = current range in tree
	// [qlo,qhi] = range to update
	// seen[i][j] = whether a cell has been visited before
	// next *[][2]int = list of new cells that were visited this round
	if qhi < loidx || qlo > hiidx {
		return
	}
	if qlo <= loidx && qhi >= hiidx {
		if s.visited[i] {
			// this is the whole reason we have a segtree, to early return when
			// no update is needed.
			return
		}
		// this range must be marked as visited
		for j := s.m + loidx; j <= s.m+hiidx; j++ {
			if s.visited[j] {
				continue
			}

			// If this cell is unvisited thus far, add to next and mark as seen
			a := s.k
			b := j - s.m
			if !s.isrow {
				a, b = b, a
			}
			if !seen[a][b] {
				*next = append(*next, [2]int{a, b})
				seen[a][b] = true
			}

			s.visited[j] = true
			for kk := j / 2; kk >= 1; kk /= 2 {
				s.visited[kk] = s.visited[kk*2] && s.visited[kk*2+1]
				if !s.visited[kk] {
					break
				}
			}
		}
		return
	}

	mid := loidx + (hiidx-loidx)/2
	s.upd(i*2, loidx, mid, qlo, qhi, seen, next)
	s.upd(i*2+1, mid+1, hiidx, qlo, qhi, seen, next)
}

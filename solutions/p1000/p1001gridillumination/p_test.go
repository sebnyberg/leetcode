package p1001gridillumination

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_gridIllumination(t *testing.T) {
	for i, tc := range []struct {
		n       int
		lamps   [][]int
		queries [][]int
		want    []int
	}{
		{
			6,
			leetcode.ParseMatrix("[[1,1]]"),
			leetcode.ParseMatrix("[[2,0],[1,0]]"),
			[]int{1, 0},
		},
		{
			6,
			leetcode.ParseMatrix("[[2,5],[4,2],[0,3],[0,5],[1,4],[4,2],[3,3],[1,0]]"),
			leetcode.ParseMatrix("[[4,3],[3,1],[5,3],[0,5],[4,4],[3,3]]"),
			[]int{1, 0, 1, 1, 0, 1},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, gridIllumination(tc.n, tc.lamps, tc.queries))
		})
	}
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	// We can't track all possible positions, but we can count the number of
	// lamps per row/col/diag/antidiag
	rows := make(map[int]int)
	cols := make(map[int]int)
	diags := make(map[int]int)
	antiDiags := make(map[int]int)
	state := make(map[[2]int]bool)
	diag := func(r, c int) int {
		return (n - 1) + (c - r)
	}
	antiDiag := func(r, c int) int {
		return (r + c)
	}
	for _, l := range lamps {
		k := [2]int{l[0], l[1]}
		if state[k] {
			continue
		}
		state[k] = true
		rows[l[0]]++
		cols[l[1]]++
		diags[diag(l[0], l[1])]++
		antiDiags[antiDiag(l[0], l[1])]++
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		r := q[0]
		c := q[1]
		if rows[r] > 0 || cols[c] > 0 || diags[diag(r, c)] > 0 || antiDiags[antiDiag(r, c)] > 0 {
			res[i] = 1
		}
		for j := max(0, r-1); j <= min(n-1, r+1); j++ {
			for k := max(0, c-1); k <= min(n-1, c+1); k++ {
				key := [2]int{j, k}
				if state[key] {
					rows[j]--
					cols[k]--
					diags[diag(j, k)]--
					antiDiags[antiDiag(j, k)]--
				}
				state[key] = false
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

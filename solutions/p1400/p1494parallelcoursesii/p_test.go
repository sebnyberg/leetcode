package p1494parallelcoursesii

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minNumberOfSemesters(t *testing.T) {
	for i, tc := range []struct {
		n         int
		relations [][]int
		k         int
		want      int
	}{
		{
			12,
			leetcode.ParseMatrix("[[11,10],[6,3],[2,5],[9,2],[4,12],[8,7],[9,5],[6,2],[7,2],[7,4],[9,3],[11,1],[4,3]]"),
			3, 4,
		},
		{
			4,
			leetcode.ParseMatrix("[[2,1],[2,4]]"),
			2, 2,
		},
		{
			5,
			leetcode.ParseMatrix("[[2,1],[3,1],[4,1],[1,5]]"),
			2, 4,
		},
		{
			4,
			leetcode.ParseMatrix("[[2,1],[3,1],[1,4]]"),
			2, 3,
		},
		{
			11,
			[][]int{},
			2, 6,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minNumberOfSemesters(tc.n, tc.relations, tc.k))
		})
	}
}

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	// At first I tried a greedy approach, but there are some edge-cases where
	// ordering based on length of courses or number of next unlocked courses
	// does not find an optimal solution. Therefore, it seems the only solution
	// that works is to explore all possible alternatives with memoization.
	prereq := make([]int, n)
	for _, r := range relations {
		a := r[0] - 1
		b := r[1] - 1
		prereq[b] |= 1 << a
	}

	// We define dp[m] to be the minimum number of semesters required to finish
	// the courses given that the i'th bit in m is set if the course has been
	// completed.
	want := (1 << n) - 1
	mem := make([]int, 1<<n)
	for i := range mem {
		mem[i] = math.MaxInt32
	}
	res := dp(mem, prereq, 0, want, k, n)
	return res
}

func dp(mem []int, prereq []int, bm, want, k, n int) int {
	if bm == want {
		return 0
	}
	if mem[bm] != math.MaxInt32 {
		return mem[bm]
	}

	// List alternatives for next courses
	cands := []int{}
	for i := range prereq {
		if bm&(1<<i) == 0 && prereq[i]&bm == prereq[i] {
			cands = append(cands, i)
		}
	}

	// Choosing k candidates is easiest to do with recursion
	wantCands := min(k, len(cands))
	mem[bm] = dp2(mem, prereq, cands, 0, wantCands, bm, want, k, n)
	return mem[bm]
}

func dp2(mem []int, prereq, cands []int, i, kk, bm, want, k, n int) int {
	if kk == 0 || i == len(cands) {
		if kk != 0 {
			return math.MaxInt32
		}
		return 1 + dp(mem, prereq, bm, want, k, n)
	}
	// Try to both skip and include this course in the current semester
	include := dp2(mem, prereq, cands, i+1, kk-1, bm|(1<<cands[i]), want, k, n)
	skip := dp2(mem, prereq, cands, i+1, kk, bm, want, k, n)
	return min(skip, include)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p1547minimumcosttocutastick

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		n    int
		cuts []int
		want int
	}{
		{7, []int{1, 3, 4, 5}, 16},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.n, tc.cuts))
		})
	}
}

func minCost(n int, cuts []int) int {
	// Should be a "hard medium" imo, not a "hard".
	// A cut divides a segment into two segments.
	// Memoize the minimum total cost for a segment..
	cuts = append(cuts, 0, n)
	m := len(cuts)
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, m)
		for j := range mem[i] {
			mem[i][j] = math.MaxInt32
		}
	}
	sort.Ints(cuts)
	res := dp(mem, cuts, 0, m-1)
	return res
}

func dp(mem [][]int, cuts []int, i, j int) int {
	if j-i == 1 {
		return 0
	}
	if mem[i][j] != math.MaxInt32 {
		return mem[i][j]
	}
	res := math.MaxInt32
	d := cuts[j] - cuts[i]
	for k := i + 1; k < j; k++ {
		x := d + dp(mem, cuts, i, k) + dp(mem, cuts, k, j)
		res = min(res, x)
	}
	mem[i][j] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

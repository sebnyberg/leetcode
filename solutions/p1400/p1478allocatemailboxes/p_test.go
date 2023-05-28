package p1478allocatemailboxes

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	for i, tc := range []struct {
		houses []int
		k      int
		want   int
	}{
		{[]int{1, 4, 8, 10, 20}, 3, 5},
		{[]int{2, 3, 5, 12, 18}, 2, 9},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minDistance(tc.houses, tc.k))
		})
	}
}

func minDistance(houses []int, k int) int {
	// Given a group of houses, the box placement that minimizes the total cost
	// is given by the median of the list of houses.
	// So what we want to do is to find the set of groups of houses that
	// minimizes the sum of the costs of the groups while covering all houses.
	// This can be done with DP
	n := len(houses)
	sort.Ints(houses)
	mem := make([][][]int, k+1)
	for kk := range mem {
		mem[kk] = make([][]int, n+1)
		for i := range mem[kk] {
			mem[kk][i] = make([]int, n+1)
			for j := range mem[kk][i] {
				mem[kk][i][j] = math.MaxInt32
			}
		}
	}
	res := dp(mem, houses, 0, n, k)
	return res
}

func dp(mem [][][]int, houses []int, i, j, k int) int {
	if mem[k][i][j] != math.MaxInt32 {
		return mem[k][i][j]
	}
	if k >= j-i {
		return 0
	}
	if k == 1 {
		m := i + (j-i)/2
		mid := houses[m]
		if (j-i)%2 == 0 {
			next := houses[m-1]
			mid = int(math.Round((float64(mid) + float64(next)) / 2))
		}
		var cost int
		for _, h := range houses[i:j] {
			cost += abs(h - mid)
		}
		mem[k][i][j] = cost
		return cost
	}
	res := math.MaxInt32
	for p := i + 1; p < j; p++ {
		left := dp(mem, houses, i, p, 1)
		right := dp(mem, houses, p, j, k-1)
		x := left + right
		res = min(res, x)
	}

	mem[k][i][j] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

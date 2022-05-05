package p1000minimumcosttomergestones

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeStones(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		k      int
		want   int
	}{
		{[]int{3, 2, 4, 1}, 2, 20},
		{[]int{3, 2, 4, 1}, 3, -1},
		{[]int{3, 5, 1, 2, 6}, 3, 25},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, mergeStones(tc.stones, tc.k))
		})
	}
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, n+1)
	for i, num := range stones {
		prefix[i+1] = prefix[i] + num
	}

	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, n+1)
		for j := range mem[i] {
			mem[i][j] = make([]int, n+1)
		}
	}
	res := dp(mem, prefix, 0, n-1, 1, k)
	if res < math.MaxInt32 {
		return res
	}
	return -1
}

func dp(mem [][][]int, prefix []int, i, j, m, k int) int {
	res := math.MaxInt32
	if mem[i][j][m] != 0 {
		return mem[i][j][m]
	}
	switch {
	case (j-i+1-m)%(k-1) != 0:
		res = math.MaxInt32
	case i == j:
		if m == 1 {
			res = 0
		} else {
			res = math.MaxInt32
		}
	case m == 1:
		res = dp(mem, prefix, i, j, k, k) + prefix[j+1] - prefix[i]
	default:
		for mid := i; mid < j; mid += k - 1 {
			res = min(res, dp(mem, prefix, i, mid, 1, k)+dp(mem, prefix, mid+1, j, m-1, k))
		}
	}
	mem[i][j][m] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

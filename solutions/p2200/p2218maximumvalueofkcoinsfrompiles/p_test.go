package p2218maximumvalueofkcoinsfrompiles

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxValueOfCoins(t *testing.T) {
	for _, tc := range []struct {
		piles [][]int
		k     int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[100],[1,700]]"),
			2, 701,
		},
		{
			leetcode.ParseMatrix("[[1,100,3],[7,8,9]]"),
			2, 101,
		},
		{
			leetcode.ParseMatrix("[[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]]"),
			7, 706,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.piles), func(t *testing.T) {
			require.Equal(t, tc.want, maxValueOfCoins(tc.piles, tc.k))
		})
	}
}

func maxValueOfCoins(piles [][]int, k int) int {
	// This is a typical max of current + max of not doing something => DP
	// With DP always try top-down first with memoization.
	mem := make(map[key]int)
	// Pre-calculate sums for piles
	sums := make([][]int, len(piles))
	for i, p := range piles {
		sums[i] = make([]int, len(p)+1)
		for j := range p {
			sums[i][j+1] = sums[i][j] + p[j]
		}
	}
	res := dp(mem, sums, 0, k, len(piles))
	return res
}

type key struct {
	i, k int
}

func dp(mem map[key]int, sums [][]int, i, k, n int) int {
	if i == n || k == 0 {
		return 0
	}
	kk := key{i, k}
	if v, exists := mem[kk]; exists {
		return v
	}
	// We can pick [0,min(k,len(sums[i])] coins from the current pile
	var res int
	for coins := 0; coins <= min(k, len(sums[i])-1); coins++ {
		// Or pick the top-most coin
		res = max(res, sums[i][coins]+dp(mem, sums, i+1, k-coins, n))
	}

	mem[kk] = res
	return mem[kk]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p2312sellingpiecesofwood

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sellingWood(t *testing.T) {
	for i, tc := range []struct {
		m, n   int
		prices [][]int
		want   int64
	}{
		{
			15, 7,
			leetcode.ParseMatrix("[[14,3,1],[2,2,1],[15,1,1],[8,1,2],[13,5,1],[6,5,1],[13,1,1],[3,3,2]]"),
			25,
		},
		{
			3, 5,
			leetcode.ParseMatrix("[[1,4,2],[2,2,7],[2,1,3]]"),
			19,
		},
		{
			4, 6,
			leetcode.ParseMatrix("[[3,2,10],[1,4,2],[4,1,3]]"),
			32,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			require.Equal(t, tc.want, sellingWood(tc.m, tc.n, tc.prices))
		})
	}
}

func sellingWood(m int, n int, prices [][]int) int64 {
	sort.Slice(prices, func(i, j int) bool {
		if prices[i][0] == prices[j][0] {
			return prices[i][1] < prices[j][1]
		}
		return prices[i][0] < prices[j][0]
	})
	var dp [201][201]int
	for i := range prices {
		hh, ww, p := prices[i][0], prices[i][1], prices[i][2]
		dp[hh][ww] = max(dp[hh][ww], p)
	}
	for h := 1; h <= m; h++ {
		for w := 1; w <= n; w++ {
			for hh := 1; hh <= h/2; hh++ {
				dp[h][w] = max(dp[h][w], dp[hh][w]+dp[h-hh][w])
			}
			for ww := 1; ww <= w/2; ww++ {
				dp[h][w] = max(dp[h][w], dp[h][ww]+dp[h][w-ww])
			}
		}
	}
	return int64(dp[m][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

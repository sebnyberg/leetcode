package p1787makexorofallsegmentsequaltozero

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minChanges(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 4, 5, 2, 1, 7, 3, 4, 7}, 3, 3},
		{[]int{1, 2, 0, 3, 0}, 1, 3},
		{[]int{1, 2, 4, 1, 2, 5, 1, 2, 6}, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minChanges(tc.nums, tc.k))
		})
	}
}

func minChanges(nums []int, k int) int {
	// To make each segment equal to zero, each segment must contain
	// the same numbers, i.e. the numbers must repeat themselves

	// Count number occurrences for each index in each segment
	indexNumCount := make([]map[int]int, k+1)
	numsAtIndex := make([]map[int]struct{}, k+1)
	for i := range indexNumCount {
		indexNumCount[i] = make(map[int]int)
		numsAtIndex[i] = make(map[int]struct{})
	}
	for i, n := range nums {
		indexNumCount[i%k][n]++
		numsAtIndex[i%k][n] = struct{}{}
	}

	// dp[i][n] contains the minimum cost to have the XOR value
	// be N at position i
	// In the end, the goal is to have the value n=0 at i=k
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, 1025)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	n := len(nums)
	minCost := 0

	for i := 0; i < k; i++ {
		nrepeat := (n + k - 1 - i) / k
		newMinCost := n + 1
		for j := 0; j < 1024; j++ {
			if i == 0 {
				dp[i][j] = nrepeat - indexNumCount[i][j]
			} else {
				for num := range numsAtIndex[i] {
					// At this position, if we wish to use "num" to reduce the cost
					// the previous position must be j^num, i.e. the previous position
					// XOR num = j
					dp[i][j] = min(dp[i][j], dp[i-1][j^num]+nrepeat-indexNumCount[i][num])
				}
				dp[i][j] = min(dp[i][j], minCost+nrepeat)
			}
			newMinCost = min(newMinCost, dp[i][j])
		}
		minCost = newMinCost
	}
	res := dp[k-1][0]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

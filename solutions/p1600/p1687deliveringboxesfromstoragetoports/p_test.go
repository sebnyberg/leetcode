package p1687deliveringboxesfromstoragetoports

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_boxDelivering(t *testing.T) {
	for i, tc := range []struct {
		boxes      [][]int
		portsCount int
		maxBoxes   int
		maxWeight  int
		want       int
	}{
		{
			leetcode.ParseMatrix("[[1,2],[3,3],[3,1],[3,1],[2,4]]"),
			3, 3, 6, 6,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, boxDelivering(tc.boxes, tc.portsCount, tc.maxBoxes, tc.maxWeight))
		})
	}
}

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	var j int     // first of the last unique box in current trip
	var k int     // first box in next trip
	var trips int // number of trips needed for current window
	for i := range boxes {
		for k < n && maxBoxes > 0 && maxWeight >= boxes[k][1] {
			maxBoxes--
			maxWeight -= boxes[k][1]
			if k == 0 || (k > 0 && boxes[k][0] != boxes[k-1][0]) {
				j = k
				trips++
			}
			k++
		}
		dp[j] = min(dp[j], dp[i]+trips)
		dp[k] = min(dp[k], dp[i]+trips+1)
		maxBoxes++
		maxWeight += boxes[i][1]
		if i < n-1 && boxes[i][0] != boxes[i+1][0] {
			trips--
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

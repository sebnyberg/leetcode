package p1824minsidewayjumps

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSideJumps(t *testing.T) {
	for _, tc := range []struct {
		obstacles []int
		want      int
	}{
		{[]int{0, 0, 3, 1, 0, 1, 0, 2, 3, 1, 0}, 2},
		{[]int{0, 1, 2, 3, 0}, 2},
		{[]int{0, 1, 1, 3, 3, 0}, 0},
		{[]int{0, 2, 1, 0, 3, 0}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.obstacles), func(t *testing.T) {
			require.Equal(t, tc.want, minSideJumps(tc.obstacles))
		})
	}
}

func minSideJumps(obstacles []int) int {
	var dp [2][4]int
	dp[0] = [4]int{math.MaxInt32, 1, 0, 1}

	for i := 1; i < len(obstacles); i++ {
		if obstacles[i] == 0 {
			continue
		}
		dp[1] = dp[0] // copy previous
		for j := 1; j <= 3; j++ {
			if j != obstacles[i] {
				continue
			}
			// Need to jump from another position
			dp[1][j] = math.MaxInt32
			for k := 1; k <= 3; k++ {
				if k == j || obstacles[i+1] == k {
					continue
				}
				dp[1][j] = min(dp[1][j], 1+dp[0][k])
			}
			continue
		}
		dp[0] = dp[1]
	}
	return min(min(dp[1][1], dp[1][2]), dp[1][3])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

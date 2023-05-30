package p1575countallpossibleroutes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countRoutes(t *testing.T) {
	for i, tc := range []struct {
		locations []int
		start     int
		finish    int
		fuel      int
		want      int
	}{
		{[]int{4, 3, 1}, 1, 0, 6, 5},
		{[]int{2, 3, 6, 8, 4}, 1, 3, 5, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countRoutes(tc.locations, tc.start, tc.finish, tc.fuel))
		})
	}
}

const mod = 1e9 + 7

func countRoutes(locations []int, start int, finish int, fuel int) int {
	// For each fuel count (200)
	// For each location (100)
	// For each other location, add count at locations[j] fuel prior to current
	n := len(locations)
	dp := make([][]int, fuel+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][start] = 1
	for k := 1; k <= fuel; k++ {
		for i, a := range locations {
			for j, b := range locations {
				if i == j {
					continue
				}
				d := abs(a - b)
				if d > k {
					continue
				}
				dp[k][i] = (dp[k][i] + dp[k-d][j]) % mod
			}
		}
	}
	var res int
	for i := range dp {
		res = (res + dp[i][finish]) % mod
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

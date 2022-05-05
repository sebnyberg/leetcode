package p1997firstdaywhereyouhavebeeninalltherooms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstDayBeenInAllRooms(t *testing.T) {
	for _, tc := range []struct {
		nextVisit []int
		want      int
	}{
		{[]int{0, 0}, 2},
		{[]int{0, 0, 2}, 6},
		{[]int{0, 1, 2, 0}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nextVisit), func(t *testing.T) {
			require.Equal(t, tc.want, firstDayBeenInAllRooms(tc.nextVisit))
		})
	}
}

const mod = 1e9 + 7

func firstDayBeenInAllRooms(nextVisit []int) int {
	// Could not finish this exercise in time because I didn't realise that
	// 0 <= nextVisit[i] <= i
	// With this constraint, the number of visits only depends on the number of
	// visits needed for prior rooms, making it easily solvable with top-down
	// dynamic programming.
	n := len(nextVisit)
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2 + mod) % mod
	}
	return dp[n-1]
}

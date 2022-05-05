package p2005subtreeremovalgamewithfibonaccitree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findGameWinner(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{4, true},
		{3, true},
		{1, false},
		{2, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findGameWinner(tc.n))
		})
	}
}

func findGameWinner(n int) bool {
	var subGraphs [101]int
	subGraphs[1] = 0 // loss
	subGraphs[2] = 1 // win
	for i := 3; i <= n; i++ {
		subGraphs[i] = (subGraphs[i-1] + 1) ^ (subGraphs[i-2] + 1)
	}
	return subGraphs[n] != 0
}

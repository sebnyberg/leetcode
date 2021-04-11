package p1802maxvaluegivenindexinboundedarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxValue(t *testing.T) {
	for _, tc := range []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		// {4, 2, 6, 2},
		// {6, 1, 10, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxValue(tc.n, tc.index, tc.maxSum))
		})
	}
}

func maxValue(n int, index int, maxSum int) int {
	// Plan
	// res := 1
	// for k := 0; ...
	// res += geometric series of left side + geometric series of right side
	// min(left side len, cur left), min(right side len, cur right)
	return 0
}

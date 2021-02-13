package p0070climbingstairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, climbStairs(tc.n))
		})
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return second
}

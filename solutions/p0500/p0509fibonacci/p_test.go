package p0509fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fib(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, fib(tc.n))
		})
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

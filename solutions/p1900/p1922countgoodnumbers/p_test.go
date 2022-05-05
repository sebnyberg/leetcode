package p1922countgoodnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countGoodNumbers(t *testing.T) {
	for _, tc := range []struct {
		n    int64
		want int
	}{
		{1, 5},
		{4, 400},
		{50, 564908303},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countGoodNumbers(tc.n))
		})
	}
}

const mod = 1_000_000_007

func countGoodNumbers(n int64) int {
	res := modPow(20, int(n)/2)
	if n%2 == 1 {
		res *= 5
	}
	return res % mod
}

func modPow(x, y int) int {
	if y == 0 {
		return 1
	}
	p := modPow(x, y/2)
	if y%2 == 1 {
		return (p * p % mod * x) % mod
	}
	return p * p % mod
}

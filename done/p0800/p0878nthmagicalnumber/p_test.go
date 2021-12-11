package p0878nthmagicalnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthMagicalNumber(t *testing.T) {
	for _, tc := range []struct {
		n    int
		a, b int
		want int
	}{
		{1, 2, 3, 2},
		{4, 2, 3, 6},
		{5, 2, 4, 10},
		{3, 6, 4, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nthMagicalNumber(tc.n, tc.a, tc.b))
		})
	}
}

const mod = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	lcm := (a * b) / gcd(a, b)
	l, r := 2, int(1e14)
	for l < r {
		mid := (l + r) / 2
		if mid/a+mid/b-mid/lcm < n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l % mod
}

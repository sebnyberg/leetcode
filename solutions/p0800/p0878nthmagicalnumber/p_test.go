package p0878nthmagicalnumber

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthMagicalNumber(t *testing.T) {
	for i, tc := range []struct {
		n    int
		a    int
		b    int
		want int
	}{
		{8, 8, 8, 64},
		{7, 5, 8, 24},
		{3, 8, 3, 8},
		{3, 6, 4, 8},
		{5, 2, 4, 10},
		{4, 2, 3, 6},
		{1, 2, 3, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, nthMagicalNumber(tc.n, tc.a, tc.b))
		})
	}
}

func nthMagicalNumber(n int, a int, b int) int {
	lcm := a * b / gcd(a, b)
	lo := 0
	hi := math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		m := mid/a + mid/b - mid/lcm
		if m < n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	const mod = 1e9 + 7
	return lo % mod
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

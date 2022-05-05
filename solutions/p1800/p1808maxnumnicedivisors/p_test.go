package p1808maxnumnicedivisors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNiceDivisors(t *testing.T) {
	for _, tc := range []struct {
		primeFactors int
		want         int
	}{
		{18, 729},
		{5, 6},
		{8, 18},
	} {
		t.Run(fmt.Sprintf("%+v", tc.primeFactors), func(t *testing.T) {
			require.Equal(t, tc.want, maxNiceDivisors(tc.primeFactors))
		})
	}
}

const mod = int(1e9) + 7

func maxNiceDivisors(primeFactors int) int {
	switch {
	case primeFactors <= 3:
		return primeFactors
	case primeFactors%3 == 0:
		return powMod(3, primeFactors/3, mod) % mod
	case primeFactors%3 == 1:
		return (powMod(3, (primeFactors-4)/3, mod) * 4) % mod
	default:
		return (2 * powMod(3, primeFactors/3, mod) % mod) % mod
	}
}

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % mod
		}
		b >>= 1
		a = (a * a) % mod
	}
	return res
}

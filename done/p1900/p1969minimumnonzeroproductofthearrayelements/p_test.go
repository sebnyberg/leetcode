package p1969minimumnonzeroproductofthearrayelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minNonZeroProduct(t *testing.T) {
	for _, tc := range []struct {
		p    int
		want int
	}{
		// {1, 1},
		// {2, 6},
		// {3, 1512},
		// {5, 202795991},
		// {34, 640964173},
		{35, 112322054},
	} {
		t.Run(fmt.Sprintf("%+v", tc.p), func(t *testing.T) {
			require.Equal(t, tc.want, minNonZeroProduct(tc.p))
		})
	}
}

const mod = 1e9 + 7

func minNonZeroProduct(p int) int {
	if p == 1 {
		return 1
	}
	if p == 2 {
		return 6
	}
	// Through just pen and paper, it seems like for any p > 2, it's possible to
	// weave 1 * (2^p)-2 * 1 * (2^p)-2 * ... * (2^p)-1
	// For example, for p = 3, 2^3 = 8. There will be three 1s, and three (2^3)-2s
	// For p = 4, there will be seven 1s and seven (2^4)-2 = 14s

	// Calculating this product for large numbers is not fast enough, so we need
	// to use modPow

	num := 1 << p
	factor := (num - 2) % mod
	exp := ((num / 2) - 1)
	res := modPow(factor, exp, mod) * ((num - 1) % mod) % mod
	return res

	// Big solution
	// num := 1 << p
	// factor := big.NewInt(int64(num - 2))
	// pow := big.NewInt(int64((num / 2) - 1))
	// bigMod := big.NewInt(int64(1e9 + 7))
	// res := factor.Exp(factor, pow, bigMod)
	// res = res.Mul(res, big.NewInt(int64(num-1)))
	// res = res.Mod(res, bigMod)
	// resInt := int(res.Int64())
	// return resInt
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}

// func modPow(a, b, mod int) int {
// 	ret := 1
// 	for ; b != 0; b /= 2 {
// 		if b%2 == 1 {
// 			ret = (ret * a) % mod
// 		}
// 		a = (a * a) % mod
// 	}
// 	return ret
// }

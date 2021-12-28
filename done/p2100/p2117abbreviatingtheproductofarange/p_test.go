package p2117abbreviatingtheproductofarange

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_abbreviateProduct(t *testing.T) {
	for _, tc := range []struct {
		left  int
		right int
		want  string
	}{
		{432, 61630, "13218...80608e15298"},
		{92882, 599690, "91828...71456e126701"},
		{410, 70833, "81384...08512e17604"},
		{256, 65535, "23510...78528e16317"},
		{2, 15, "1307674368e3"},
		{1000000, 1000000, "1e6"},
		{8, 18, "12703...22432e2"},
		{1, 4, "24e0"},
		{2, 11, "399168e2"},
		{999998, 1000000, "99999...00002e6"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, abbreviateProduct(tc.left, tc.right))
		})
	}
}

// func abbreviateProduct(left int, right int) string {
// 	pref, suf := 1, 1
// 	var zeroes int
// 	var dosplit bool
// 	var prefPrecision int = 1e12
// 	var sufPrecision int = 1e12
// 	for i := left; i <= right; i++ {
// 		suf *= i
// 		pref *= i
// 		for pref >= prefPrecision {
// 			pref /= 10
// 		}
// 		for suf%10 == 0 {
// 			zeroes++
// 			suf /= 10
// 		}
// 		if suf >= 1e10 {
// 			dosplit = true
// 		}
// 		suf %= sufPrecision
// 	}
// 	for pref >= 1e5 {
// 		pref /= 10
// 	}
// 	if dosplit {
// 		suf %= 1e5
// 		return fmt.Sprintf("%d...%05de%d", pref, suf, zeroes)
// 	}
// 	return fmt.Sprintf("%de%d", suf, zeroes)
// }

func abbreviateProduct(left int, right int) string {
	// Remove trailing zeroes on each multiplication
	// Split the result into prefix and suffix
	// The largest possible suffix after a multiplication is 1e6 * 1e6 = 1e12
	// Store trailing zeroes
	// Remove from prefix until there are only 5 digits
	pref := big.NewInt(1)
	suf := big.NewInt(1)

	// Very large max value
	prefDiv := new(big.Int)
	prefDiv.Exp(big.NewInt(10), big.NewInt(15), nil)
	maxPref := new(big.Int)
	maxPref.Exp(big.NewInt(10), big.NewInt(30), nil)

	// And a large divisor
	bigMod := new(big.Int)
	bigMod.Exp(big.NewInt(10), big.NewInt(10), nil)
	ten := big.NewInt(10)

	var zeroes int
	// chunkSize := int64(5)
	for i := int64(left); i <= int64(right); i++ {
		// r := min(int64(right), i+chunkSize-1)

		// rng := big.NewInt(0).MulRange(i, r)
		pref = pref.Mul(pref, big.NewInt(i))
		suf = suf.Mul(suf, big.NewInt(i))

		if pref.Cmp(maxPref) == 1 {
			pref = pref.Quo(pref, prefDiv)
		}
		for big.NewInt(0).Rem(suf, ten).Uint64() == 0 {
			suf.Quo(suf, ten)
			zeroes++
		}
		suf = suf.Rem(suf, bigMod)
	}

	prefStr := pref.String()
	sufStr := suf.String()
	if len(sufStr) > 10 {
		return fmt.Sprintf("%s...%se%d", prefStr[:5], sufStr[len(sufStr)-5:], zeroes)
	}
	return fmt.Sprintf("%se%d", sufStr, zeroes)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

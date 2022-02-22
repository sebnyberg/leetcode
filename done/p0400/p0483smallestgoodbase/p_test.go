package p0483smallestgoodbase

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestGoodBase(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want string
	}{
		{"13", "3"},
		{"4681", "8"},
		{"1000000000000000000", "999999999999999999"},
		{"16035713712910627", "502"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, smallestGoodBase(tc.n))
		})
	}
}

func smallestGoodBase(n string) string {
	// The solution can be written on the form
	// z^m + z^(m-1) + ... + z^1 + 1 = n
	//
	// This is equivalent to the geometric series
	// (1 - z^(m+1)) / (1 - z) = n
	//
	// Note that the solution must be divisible by (1-z)
	//
	// The greatest m is for base 2 (z = 2), so max m is given by:
	// log_2(2^m) >= log_2(n)
	// <=>
	// m*log_2(2) >= log_2(n)
	// <=>
	// m >= log_2(n)
	//
	fn, _ := strconv.ParseFloat(n, 64)
	intN, _ := strconv.ParseInt(n, 10, 64)
	bigN := big.NewInt(intN)
	maxM := int(math.Ceil(math.Log2(fn)))
	for m := int64(maxM); m >= 1; m-- {
		// So given z^m + z^(m-1) + ... + z + 1 = n
		// What z's are feasible?
		// z^m < n
		// Also, (z+1)^m > n
		// This gives z^m < n < (z+1)^m
		// => z < n^(1/m) < (z+1)
		//
		// The integer root n^(1/m) is a maybe floored version of the actual result,
		// in which case it can be the solution z. Otherwise, it's not z. In either
		// case, it's impossible for n^(1/m) to end up being (z+1).
		//
		// And so, the only real candidate is n^(1/m)
		cand := int64(math.Pow(fn, 1/float64(m)))
		nom := new(big.Int).Exp(big.NewInt(cand), big.NewInt(m+1), nil)
		nom.Sub(nom, big.NewInt(1))
		divisor := new(big.Int).Sub(big.NewInt(cand), big.NewInt(1))
		if divisor.String() == "0" {
			continue
		}
		res := new(big.Int).Div(nom, divisor)
		if res.String() == bigN.String() {
			return fmt.Sprint(cand)
		}
	}
	return fmt.Sprint(intN - 1)
}

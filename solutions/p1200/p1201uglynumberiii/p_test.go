package p1201uglynumberiii

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthUglyNumber(t *testing.T) {
	for i, tc := range []struct {
		n       int
		a, b, c int
		want    int
	}{
		{5, 2, 3, 3, 8},
		{4, 2, 3, 4, 6},
		{3, 2, 3, 5, 4},
		{5, 2, 11, 13, 10},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, nthUglyNumber(tc.n, tc.a, tc.b, tc.c))
		})
	}
}

func nthUglyNumber(n int, a int, b int, c int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Calculate LCMs
	// Note the constraint that a*b*c <= 10^18, otherwise there would be 64-bit
	// overflow here.
	ab := (a * b) / gcd(a, b)
	bc := (b * c) / gcd(b, c)
	ac := (a * c) / gcd(a, c)
	abc := ab * c / gcd(ab, c)

	// The easiest way to check the value is to guess.
	// check returns true if there exists >= n ugly numbers at or prior to x
	check := func(x int) bool {
		// count ugly numbers prior to this number due to a
		count := x / a

		// count ugly numbers due to b, disregarding numbers that are shared
		// with a
		count += x/b - x/ab

		// count ugly numbers due to c, disregarding overlap with a and b and
		// adding overlap with all three (would be double-removed otherwise)
		count += x/c - x/ac - x/bc + x/abc

		return count >= n
	}

	lo, hi := 0, math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

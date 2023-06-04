package p2338countthenumberofidealarrays

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_idealArrays(t *testing.T) {
	for i, tc := range []struct {
		n        int
		maxValue int
		want     int
	}{
		{10, 1, 1},
		{2, 5, 10},
		{5, 3, 11},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, idealArrays(tc.n, tc.maxValue))
		})
	}
}

const mod = 1e9 + 7

func idealArrays(n int, maxValue int) int {
	// This is a pretty nasty problem.
	//
	// Given a number val, where 1 <= val <= maxValue
	//
	// The value consists of a number of primes and counts of primes.
	// Any valid sequence that ends with that number contains the prime factors
	// as multipliers along the sequence. For example, 12 has prime factors {2,
	// 2, 3}. This means that 2, 2, and 3 must be placed somewhere in the
	// sequence and they may overlap!
	//
	// For example, with n = 5, we may pile all three factors in the first spot:
	//
	//       factors:  [ {2,2,3}, {}, {}, {}, {} ]
	// This gives us:  [    12,   12, 12, 12, 12 ]
	//
	// Or we could have: [ {3}, {2}, {}, {2}, {} ]
	//      which gives: [  3,   6,   6, 12,  12 ]
	//
	// Therefore, to find the number of possible valid sequences ending in a
	// certain value, we calculate nCk(n, k), but we must also consider that
	// multiple factors are allowed to overlap.

	// Generate binomial coefficients using Pascals triangle
	maxK := int(math.Log2(float64(maxValue)))
	m := n + maxK
	bins := make([][]int, m+1)
	for i := range bins {
		bins[i] = make([]int, maxK+1)
	}
	bins[0][0] = 1
	for i := 1; i < len(bins); i++ {
		bins[i][0] = 1
		for k := 1; k <= min(maxK, i); k++ {
			bins[i][k] = (bins[i-1][k-1] + bins[i-1][k]) % mod
		}
		if i < len(bins[i]) {
			bins[i][i] = 1
		}
	}

	factorize := func(x int) map[int]int {
		if x == 1 {
			return map[int]int{}
		}
		res := make(map[int]int)
		for y := 2; y*y <= x; y++ {
			for x%y == 0 {
				res[y]++
				x /= y
			}
		}
		if x > 1 {
			res[x] = 1
		}
		return res
	}

	var res int
	for x := 1; x <= maxValue; x++ {
		primes := factorize(x)
		cur := 1
		for _, cnt := range primes {
			ways := bins[n+cnt-1][cnt]
			cur = (cur * ways) % mod
		}
		res = (res + cur) % mod
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

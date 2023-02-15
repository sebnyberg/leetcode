package p1175primearrangements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numPrimeArrangements(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{5, 12},
		{100, 12},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numPrimeArrangements(tc.n))
		})
	}
}

func numPrimeArrangements(n int) int {
	// This is not an easy problem. It requires some understanding of
	// combinatorics.
	//
	// In any case. The idea is that primes can only be in prime positions, and
	// non-primes only in non-prime positions. This restricts the options for
	// choosing a combination.
	//
	// Let's take 5 as an example. The primes are {2,3,5}. For the first prime
	// index, we can choose between all three, givin us three options. Then two,
	// then 1. I.e. 3! For non-prime positions, we can choose between {1,4},
	// i.e. 2!
	//
	// So given 'p' primes, the number of combinations are (n-p)!p!
	//
	notprime := make([]bool, n+1)
	notprime[1] = true
	for x := 2; x <= n; x++ {
		if notprime[x] {
			continue
		}
		for y := x + x; y <= n; y += x {
			notprime[y] = true
		}
	}
	var primes int
	for x := 1; x <= n; x++ {
		if !notprime[x] {
			primes++
		}
	}
	primeCombs := 1
	const mod = 1e9 + 7
	for x := primes; x > 1; x-- {
		primeCombs = (primeCombs * x) % mod
	}
	nonPrimeCombs := 1
	for x := n - primes; x > 1; x-- {
		nonPrimeCombs = (nonPrimeCombs * x) % mod
	}
	return (primeCombs * nonPrimeCombs) % mod
}

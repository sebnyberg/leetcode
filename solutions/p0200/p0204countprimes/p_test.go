package p0204countprimes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPrimes(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{10, 4},
		{0, 0},
		{1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countPrimes(tc.n))
		})
	}
}

func countPrimes(n int) int {
	sieve := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if sieve[i] {
			continue
		}
		for j := i; j < n; j += i {
			sieve[j] = true
		}
		res++
	}

	return res
}

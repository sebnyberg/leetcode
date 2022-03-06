package p0479largestpalindromeproduct

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 987},
		{1, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, largestPalindrome(tc.n))
		})
	}
}

const mod = 1337

func largestPalindrome(n int) int {
	// The result is either of length 2*n or 2*n-1
	max := pow(10, n) - 1
	min := max / 10

	for m := max; m > min; m-- {
		left := m
		var right int
		for x := m; x > 0; x /= 10 {
			right *= 10
			right += x % 10
			left *= 10
		}
		cand := left + right

		for factor := max; factor > min; factor-- {
			d := cand / factor
			if d > factor {
				// Continuing would only try the same numbers once again
				break
			}
			if cand%factor == 0 {
				return cand % mod
			}
		}
	}
	return 9
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

func pow(a, b int) int {
	res := a
	for i := 1; i < b; i++ {
		res *= a
	}
	return res
}

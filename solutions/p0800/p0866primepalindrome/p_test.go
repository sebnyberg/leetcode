package p0866primepalindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_primePalindrome(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{13, 101},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, primePalindrome(tc.n))
		})
	}
}

func primePalindrome(n int) int {
	// There aren't that many palindromes so we can try them all.

	numlen := func(x int) int {
		var res int
		for x > 0 {
			x /= 10
			res++
		}
		return res
	}
	left := 1
	if n == 1 {
		left++
	}
	palin := func(x int, div int) int {
		res := x
		x /= div
		for x > 0 {
			res *= 10
			res += x % 10
			x /= 10
		}
		return res
	}
	isprime := func(x int) bool {
		if x == 1 {
			return false
		}
		for a := 2; a*a <= x; a++ {
			if x%a == 0 {
				return false
			}
		}
		return true
	}
	div := 1
	if numlen(n)&1 == 1 {
		div = 10
	}
	for {
		p := palin(left, div)
		if numlen(left+1) > numlen(left) {
			if div == 10 {
				div = 1
				left = (left + 1) / 10
			} else {
				div = 10
				left++
			}
		} else {
			left++
		}
		if p < n {
			continue
		}
		if isprime(p) {
			return p
		}
	}
}

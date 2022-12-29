package p2514countanagrams

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countAnagrams(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"b okzojaporykbmq tybq zrztwlolvcyumcsq jjuowpp", 210324488},
		{"too hot", 18},
		{"aa", 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countAnagrams(tc.s))
		})
	}
}

const mod = 1e9 + 7

func countAnagrams(s string) int {
	// This is a 100% math problem. Kinda annoying.
	//
	// The number of distinct permutations for a given word is equal to the
	// total number of permutations divided by the number of ways duplicate
	// letters can swap positions.
	//
	// The difficult part is knowing Euler's theorem for modular inverses,
	// i.e. that
	//
	//     a/b mod p
	// <=> a*inv(b) mod p
	//
	// where inv(b) = b^(p-2) % p
	//
	// This is only true when b and p are co-prime, which they are when p is a
	// prime.
	//
	// Note that the modular inverse only exists for Galois finite fields (of
	// power-of-prime orders) as proved by the 20-year-old Evariste Gaiois. What
	// are you doing with your life?
	//
	countChars := func(t string) [26]int {
		var count [26]int
		for i := range t {
			count[t[i]-'a']++
		}
		return count
	}
	res := 1
	for _, w := range strings.Fields(s) {
		count := countChars(w)
		var tot int
		dup := 1
		for _, x := range count {
			tot += x
			if x > 1 {
				dup = (dup * fac(x)) % mod
			}
		}
		inv := modInverse(dup, mod)
		wordRes := (fac(tot) * inv) % mod
		res = (res * wordRes) % mod
	}
	return res % mod
}

func fac(x int) int {
	res := 1
	for x >= 1 {
		res = (res * x) % mod
		x--
	}
	return res
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
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

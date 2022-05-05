package p2002maximumproductofthelengthoftwopalindromicsubsequences

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	res := isPalindrome("leetcode",
		(1<<1)|(1<<3)|(1<<7),
	)
	require.True(t, res)
}

func Test_maxProduct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"leetcodecom", 9},
		{"bb", 1},
		{"accbcaxxcxx", 25},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, maxProduct(tc.s))
		})
	}
}

func maxProduct(s string) int {
	// There are 2^12 (4k) different possible palindrome subsequences in s.
	// Finding all possible subsequences is relatively easy.
	// Once subsequences are known, the goal is to match two palindromes together
	// so that they maximize the product of their lengths.
	// Doing a for each palindrome pair comparison with bitwise or results in
	// O(n^2) => 1.6*10^7 operations. Slow but may pass without TLE.
	var f palindromeFinder
	f.palindromes = make(map[uint]struct{})
	f.findPalindromes(s, 0, 0, len(s))
	allPalindromes := make([]uint, 0, len(f.palindromes))
	for bm := range f.palindromes {
		allPalindromes = append(allPalindromes, bm)
	}
	res := 1
	for i := 0; i < len(allPalindromes)-1; i++ {
		for j := i + 1; j < len(allPalindromes); j++ {
			if allPalindromes[i]&allPalindromes[j] == 0 {
				first := bits.OnesCount(uint(allPalindromes[i]))
				second := bits.OnesCount(uint(allPalindromes[j]))
				res = max(res, first*second)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type palindromeFinder struct {
	palindromes map[uint]struct{}
}

func (f *palindromeFinder) findPalindromes(s string, bm uint, idx, n int) {
	if idx == n {
		if bm != 0 && isPalindrome(s, bm) {
			f.palindromes[bm] = struct{}{}
		}
		return
	}
	f.findPalindromes(s, bm, idx+1, n)
	f.findPalindromes(s, bm|(1<<idx), idx+1, n)
}

func isPalindrome(s string, bm uint) bool {
	palinStr := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if bm&(1<<i) > 0 {
			palinStr = append(palinStr, s[i])
		}
	}
	for l, r := 0, len(palinStr)-1; l < r; l, r = l+1, r-1 {
		if palinStr[l] != palinStr[r] {
			return false
		}
	}
	return true
}

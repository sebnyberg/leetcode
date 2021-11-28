package p1960maxproductofthelengthoftwopalindromicsubstrings

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_maxProduct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int64
	}{
		{"ggbswiymmlevedhkbdhntnhdbkhdevelmmyiwsbgg", 45},
		{"aaabaaaba", 15},
		{"ababbb", 9},
		{"zaaaxbbby", 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, maxProduct(tc.s))
		})
	}
}

func maxProduct(s string) int64 {
	// Manacher's Algorithm (confined to odd-only results)
	// The goal of this algorithm is to re-use sub-palindromes confined within
	// larger palindromes.
	//
	// For example, when visiting i=3 for:
	// 'aaabaaa'
	//  0123456
	// A palindrome is found with the size 7.
	// Once i=5, then there will be another palindrome of size 3 ('aaa')
	manachers := func(ss string, n int) []int {
		m := make([]int, n)
		maxLeft := make([]int, n)
		for i := range maxLeft {
			maxLeft[i] = 1
		}
		for i, l, r := 0, 0, -1; i < n; i++ {
			k := 1
			if i <= r {
				k = min(m[l+r-i], r-i+1)
			}
			for i-k >= 0 && i+k < n && ss[i-k] == ss[i+k] {
				// Not part of Manacher's Algo: keep track of the maximum palindrome
				// found to the left of where the current palindrome ends:
				maxLeft[i+k] = 2*k + 1
				k++
			}
			m[i] = k
			if i+k > r {
				l = i - k + 1
				r = i + k - 1
			}
		}
		// Fill graps in the prefix max with longest palindrome.
		for i := 1; i < n; i++ {
			maxLeft[i] = max(maxLeft[i], maxLeft[i-1])
		}
		return maxLeft
	}

	n := utf8.RuneCountInString(s)
	maxLeft := manachers(s, n)

	// To get the maximum palindrome to the right of the string, reverse the input
	// and the result to get maxRight.
	rev := []byte(s)
	for l, r := 0, len(rev)-1; l < r; l, r = l+1, r-1 {
		rev[l], rev[r] = rev[r], rev[l]
	}
	maxRight := manachers(string(rev), n)
	for l, r := 0, len(maxRight)-1; l < r; l, r = l+1, r-1 {
		maxRight[l], maxRight[r] = maxRight[r], maxRight[l]
	}

	// Find maximum product
	res := 1
	for i := 0; i < len(maxLeft)-1; i++ {
		res = max(res, maxLeft[i]*maxRight[i+1])
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p3333findtheoriginaltypedstringsii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_possibleStringCount(t *testing.T) {
	for _, tc := range []struct {
		word string
		k    int
		want int
	}{
		{"aabbccdd", 7, 5},
		{"aabbccdd", 8, 1},
		{"aabbccdd", 4, 16},
		{"abc", 3, 1},
		{"abc", 4, 0},
		{"aaa", 1, 3},
		{"aaa", 2, 2},
		{"aaa", 3, 1},
		{"aaa", 4, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, possibleStringCount(tc.word, tc.k))
		})
	}
}

const mod = 1e9 + 7

func possibleStringCount(word string, k int) int {
	n := len(word)

	// Find groups of consecutive identical characters
	groups := []int{}
	count := 1
	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			groups = append(groups, count)
			count = 1
		}
	}
	groups = append(groups, count)

	m := len(groups) // number of unique character groups

	// Total possible strings (without length constraint)
	total := 1
	for _, g := range groups {
		total = (total * g) % mod
	}

	// If minimum possible length (m) is already >= k, return total
	if m >= k {
		return total
	}

	// Use optimized 1D DP to count invalid strings
	// Invalid = strings with length < k
	// Transform: each group contributes 1 base + (0 to groupSize-1) extra
	invalid := 0

	// dp[j] = ways to achieve j extra characters (beyond the minimum m)
	dp := make([]int, k-m)
	dp[0] = 1

	for _, groupSize := range groups {
		// Process in reverse to avoid using updated values in same iteration
		for j := k - m - 1; j >= 0; j-- {
			if dp[j] > 0 {
				ways := dp[j]
				// Add extra characters from this group (1 to groupSize-1 extra)
				for extra := 1; extra < groupSize && j+extra < k-m; extra++ {
					dp[j+extra] = (dp[j+extra] + ways) % mod
				}
			}
		}
	}

	// Sum all invalid combinations
	invalid = 0
	for j := 0; j < k-m; j++ {
		invalid = (invalid + dp[j]) % mod
	}

	// Result = total - invalid
	return (total - invalid + mod) % mod
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

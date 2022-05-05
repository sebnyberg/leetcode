package p0471encodestringwithshortestlength

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_encode(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"abbbabbbcabbbabbbc", "2[2[abbb]c]"},
		{"aaa", "aaa"},
		{"aaaaa", "5[a]"},
		{"aaaaaaaaaa", "10[a]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, len(tc.want), len(encode(tc.s)))
		})
	}
}

func encode(s string) string {
	mem := make(map[string]string)
	res := encodeCached(mem, s)
	return res
}

func encodeCached(mem map[string]string, s string) string {
	if _, exists := mem[s]; exists {
		return mem[s]
	}
	if len(s) == 0 {
		return ""
	}
	if len(s) <= 4 {
		return s
	}
	res := s
	for k := len(s) / 2; k > 0; k-- {
		// Attempt to find patterns which repeat and cover s
		if len(s)%k != 0 {
			continue
		}
		// Apply pattern from [0,k) across s to see how many times it repeats
		count := 1
		for i := k; i+k <= len(s); i += k {
			if s[i:i+k] != s[:k] {
				break
			}
			count++
		}
		if count*k != len(s) {
			continue
		}
		encodedPattern := encodeCached(mem, s[:k])
		cand := fmt.Sprintf("%v[%s]", count, encodedPattern)
		if len(cand) < len(res) {
			res = cand
		}
	}

	// Either the pattern found previously is the optimum, or it is possible to
	// partition the string in such a way that there are better encodings
	// Try each such partition
	for i := 1; i < len(s); i++ {
		left := encodeCached(mem, s[:i])
		right := encodeCached(mem, s[i:])
		if len(left)+len(right) < len(res) {
			res = left + right
		}
	}

	mem[s] = res
	return string(res)
}

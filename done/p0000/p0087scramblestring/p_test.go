package p0087scramblestring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isScramble(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want bool
	}{
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
		{"a", "a", true},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s1, tc.s2), func(t *testing.T) {
			require.Equal(t, tc.want, isScramble(tc.s1, tc.s2))
		})
	}
}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	if n != len(s2) {
		return false
	}

	// Finding a way to scramble s1 into s2 involves finding a cut
	// which partitions s1 into two parts: [l,r] such that a count
	// of the runes in [l,r] in s1 OR [r,l] matches the corresponding
	// sections in s2

	// Count number of runes and match it
	var runeCount [26]byte
	for i := range s1 {
		runeCount[s1[i]-'a']++
		runeCount[s2[i]-'a']--
	}
	for _, cnt := range runeCount {
		if cnt != 0 {
			return false
		}
	}

	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}

	return false
}

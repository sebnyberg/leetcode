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
	// Finding a way to scramble s1 into s2 involves finding a cut
	// which partitions s1 into two parts: [l,r] such that a count
	// of the runes in [l,r] in s1 OR [r,l] matches the corresponding
	// sections in s2

	pre1 := make([][26]int, len(s1)+1)
	pre2 := make([][26]int, len(s2)+1)

	for i := range s1 {
		pre1[i+1] = pre1[i]
		pre1[i+1][s1[i]-'a']++
	}

	for i := range s2 {
		pre2[i+1] = pre2[i]
		pre2[i+1][s2[i]-'a']++
	}
	diff := func(a, b [26]int) [26]int {
		var res [26]int
		for i := range a {
			res[i] = b[i] - a[i]
		}
		return res
	}

	mem := make(map[[2]string]bool)

	var cmp func(s1, s2 string, p1, p2 [][26]int) bool
	cmp = func(s1, s2 string, p1, p2 [][26]int) bool {
		if s1 == s2 {
			return true
		}
		k := [2]string{s1, s2}
		if v, exists := mem[k]; exists {
			return v
		}
		n := len(s1)
		if n != len(s2) {
			return false
		}
		c1 := diff(p1[n], p1[0])
		c2 := diff(p2[n], p2[0])
		if c1 != c2 {
			return false
		}
		for i := 1; i < n; i++ {
			if cmp(s1[:i], s2[:i], p1[:i+1], p2[:i+1]) &&
				cmp(s1[i:], s2[i:], p1[i:], p2[i:]) {
				mem[k] = true
				return true
			}
			if cmp(s1[:i], s2[n-i:], p1[:i+1], p2[n-i:]) &&
				cmp(s1[i:], s2[:n-i], p1[i:], p2[:n-i+1]) {
				mem[k] = true
				return true
			}
		}
		mem[k] = false
		return false
	}

	return cmp(s1, s2, pre1, pre2)
}

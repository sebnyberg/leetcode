package p0567permutationinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkInclusion(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, checkInclusion(tc.s1, tc.s2))
		})
	}
}

func checkInclusion(s1 string, s2 string) bool {
	var s1Count [26]int
	for _, ch := range s1 {
		s1Count[ch-'a']++
	}
	var s2Count [26]int
	for i, ch := range s2 {
		s2Count[ch-'a']++
		if i >= len(s1) {
			s2Count[s2[i-len(s1)]-'a']--
		}
		if s1Count == s2Count {
			return true
		}
	}
	return false
}

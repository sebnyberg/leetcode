package p0266palindromeperm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPermutePalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"carerac", true},
		{"code", false},
		{"aab", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, canPermutePalindrome(tc.s))
		})
	}
}

func canPermutePalindrome(s string) bool {
	var chCount [26]int
	for _, ch := range s {
		chCount[ch-'a']++
	}

	var hasOdd bool
	for _, count := range chCount {
		if count%2 == 1 {
			if hasOdd {
				return false
			}
			hasOdd = true
			continue
		}
	}
	return true
}

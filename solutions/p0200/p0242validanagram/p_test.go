package p0242validanagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isAnagram(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.s, tc.t), func(t *testing.T) {
			require.Equal(t, tc.want, isAnagram(tc.s, tc.t))
		})
	}
}

func isAnagram(s string, t string) bool {
	countChars := func(ss string) [26]int {
		var count [26]int
		for _, ch := range ss {
			count[ch-'a']++
		}
		return count
	}
	a := countChars(s)
	b := countChars(t)
	return a == b
}

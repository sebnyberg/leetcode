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
	sstrs := make(map[rune]int)
	var n int
	for _, r := range s {
		sstrs[r]++
		n++
	}
	for _, r := range t {
		if sstrs[r] <= 0 {
			return false
		}
		sstrs[r]--
		n--
	}
	return n == 0
}

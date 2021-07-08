package p0383ransomnote

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canConstruct(t *testing.T) {
	for _, tc := range []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ransomNote), func(t *testing.T) {
			require.Equal(t, tc.want, canConstruct(tc.ransomNote, tc.magazine))
		})
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	var count [26]int
	for _, ch := range ransomNote {
		count[ch-'a']++
	}
	for _, ch := range magazine {
		count[ch-'a']--
	}
	for _, n := range count {
		if n > 0 {
			return false
		}
	}
	return true
}

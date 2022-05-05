package p0389findthedifference

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheDifference(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want byte
	}{
		{"abcd", "abcde", 'e'},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findTheDifference(tc.s, tc.t))
		})
	}
}

func findTheDifference(s string, t string) byte {
	var count [26]int
	for _, ch := range t {
		count[ch-'a']++
	}
	for _, ch := range s {
		count[ch-'a']--
	}
	for ch, cnt := range count {
		if cnt > 0 {
			return byte(ch + 'a')
		}
	}
	return '0'
}

package p0392issubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isSubsequence(tc.s, tc.t))
		})
	}
}

func isSubsequence(s string, t string) bool {
	var pos int
	for i := range t {
		if pos == len(s) {
			return true
		}
		if t[i] == s[pos] {
			pos++
		}
	}
	return pos == len(s)
}

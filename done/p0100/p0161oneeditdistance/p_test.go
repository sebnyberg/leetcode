package p0161oneeditdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isOneEditDistance(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"cb", "ab", true},
		{"a", "ac", true},
		{"ab", "acb", true},
		{"", "", false},
		{"a", "", true},
		{"", "A", true},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.t), func(t *testing.T) {
			require.Equal(t, tc.want, isOneEditDistance(tc.s, tc.t))
		})
	}
}

func isOneEditDistance(s string, t string) bool {
	if len(t) > len(s) {
		s, t = t, s
	}
	// t is the shorter string
	var i, j int
	edits := 0
	for i = range s {
		switch {
		case j >= len(t) || s[i] != t[j]:
			edits++
			if edits == 2 {
				return false
			}
			if len(s) == len(t) {
				j++
			}
		case s[i] == t[j]:
			j++
		}
	}
	return edits == 1
}

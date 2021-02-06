package p0044wildcardmatching

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isMatch(t *testing.T) {
	for _, tc := range []struct {
		s    string
		p    string
		want bool
	}{
		{"abcabczzzde", "*abc???de*", true},
		{"zacabz", "*a?b*", false},
		{"aab", "c*a*b", false},
		{"acdcb", "a*c?b", false},
		{"adceb", "*a*b", true},
		{"cb", "?a", false},
		{"aa", "*", true},
		{"aa", "a", false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.p), func(t *testing.T) {
			require.Equal(t, tc.want, isMatch(tc.s, tc.p))
		})
	}
}

func isMatch(s string, p string) bool {
	d := make([][]bool, len(p)+1)
	for i := range d {
		d[i] = make([]bool, len(s)+1)
	}
	for i := 0; i < len(p) && p[i] == '*'; i++ {
		d[i+1][0] = true
	}
	d[0][0] = true

	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			switch p[i-1] {
			case s[j-1]:
				d[i][j] = d[i-1][j-1]
			case '*':
				if d[i-1][j-1] || d[i-1][j] {
					for j <= len(s) {
						d[i][j] = true
						j++
					}
				}
			case '?':
				d[i][j] = d[i-1][j-1]
			default:
			}
		}
	}
	return d[len(p)][len(s)]
}

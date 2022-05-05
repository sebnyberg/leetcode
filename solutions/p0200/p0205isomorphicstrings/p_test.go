package p0205isomorphicstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIsomorphic(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"foo", "bar", false},
		{"egg", "add", true},
		{"paper", "title", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isIsomorphic(tc.s, tc.t))
		})
	}
}

func isIsomorphic(s string, t string) bool {
	ss := make(map[byte]int)
	tt := make(map[byte]int)
	for i := range s {
		ssVal, ssExists := ss[s[i]]
		ttVal, ttExists := tt[t[i]]
		switch {
		case ssExists && !ttExists,
			!ssExists && ttExists:
			return false
		case !ssExists && !ttExists:
			ss[s[i]] = i
			tt[t[i]] = i
		default:
			if ssVal != ttVal {
				return false
			}
		}
	}
	return true
}

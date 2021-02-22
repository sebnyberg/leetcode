package p0524longestwordindict

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLongestWord(t *testing.T) {
	for _, tc := range []struct {
		s    string
		d    []string
		want string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
		{"bab", []string{"ba", "ab", "a", "b"}, "ab"},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.s, tc.d), func(t *testing.T) {
			require.Equal(t, tc.want, findLongestWord(tc.s, tc.d))
		})
	}
}

func findLongestWord(s string, d []string) string {
	// Sort strings
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) > len(d[j]) {
			return true
		}
		if len(d[i]) < len(d[j]) {
			return false
		}
		// equal length, sort lexicographically
		for k := range d[i] {
			if d[i][k] == d[j][k] {
				continue
			}
			if d[i][k] < d[j][k] {
				return true
			}
			return false
		}
		return false
	})

	for i := 0; i < len(d); i++ {
		var k int
		ss := d[i]
		for i := range s {
			if k == len(ss) {
				break
			}
			if s[i] == ss[k] {
				k++
			}
		}
		if k == len(ss) {
			return ss
		}
	}

	return ""
}

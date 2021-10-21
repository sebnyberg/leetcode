package p1957deletecharacterstomakefancystring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeFancyString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"leeetcode", "leetcode"},
		{"aaabaaaa", "aabaa"},
		{"aab", "aab"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, makeFancyString(tc.s))
		})
	}
}

func makeFancyString(s string) string {
	res := make([]byte, 0, len(s))
	res = append(res, s[0])
	i := 0
	count := 1
	for j := 1; j < len(s); j++ {
		if res[i] == s[j] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			res = append(res, s[j])
			i++
		}
	}
	return string(res)
}

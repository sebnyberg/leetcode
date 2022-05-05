package p0616addboldtaginstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBoldTag(t *testing.T) {
	for _, tc := range []struct {
		s     string
		words []string
		want  string
	}{
		{"abcxyz123", []string{"abc", "123"}, "<b>abc</b>xyz<b>123</b>"},
		{"aaabbcc", []string{"aaa", "aab", "bc"}, "<b>aaabbc</b>c"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, addBoldTag(tc.s, tc.words))
		})
	}
}

func addBoldTag(s string, words []string) string {
	n := len(s)
	var bolded [1001]bool
	for _, w := range words {
		for i := 0; i < n-len(w)+1; i++ {
			if s[i:i+len(w)] == w {
				for j := 0; j < len(w); j++ {
					bolded[i+j] = true
				}
			}
		}
	}
	bold := false
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if bolded[i] && !bold {
			res = append(res, "<b>"...)
			bold = true
		}
		if !bolded[i] && bold {
			res = append(res, "</b>"...)
			bold = false
		}
		res = append(res, s[i])
	}
	if bold {
		res = append(res, "</b>"...)
	}
	return string(res)
}

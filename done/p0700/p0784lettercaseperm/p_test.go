package p0784lettercaseperm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_letterCasePermutation(t *testing.T) {
	for _, tc := range []struct {
		S    string
		want []string
	}{
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"3z4", []string{"3z4", "3Z4"}},
		{"12345", []string{"12345"}},
		{"0", []string{"0"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.S), func(t *testing.T) {
			require.Equal(t, tc.want, letterCasePermutation(tc.S))
		})
	}
}

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	helper(0, S, []rune{}, &res)
	return res
}

const capitalizationDistance int = 'a' - 'A'

func helper(i int, s string, prefix []rune, res *[]string) {
	if i == len(s) {
		*res = append(*res, string(prefix))
		return
	}
	if s[i] >= 'A' && s[i] <= 'Z' {
		prefixCpy := make([]rune, len(prefix))
		copy(prefixCpy, prefix)
		prefix = append(prefix, rune(s[i]))
		prefixCpy = append(prefixCpy, rune(s[i]+byte(capitalizationDistance)))
		helper(i+1, s, prefix, res)
		helper(i+1, s, prefixCpy, res)
	} else if s[i] >= 'a' && s[i] <= 'z' {
		prefixCpy := make([]rune, len(prefix))
		copy(prefixCpy, prefix)
		prefix = append(prefix, rune(s[i]))
		prefixCpy = append(prefixCpy, rune(s[i]-byte(capitalizationDistance)))
		helper(i+1, s, prefix, res)
		helper(i+1, s, prefixCpy, res)
	} else {
		prefix = append(prefix, rune(s[i]))
		helper(i+1, s, prefix, res)
	}
}

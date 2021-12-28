package p1153stringtransformsintoanotherstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canConvert(t *testing.T) {
	for _, tc := range []struct {
		str1 string
		str2 string
		want bool
	}{
		{"aa", "ab", false},
		{"abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza", false},
		{"abcdefghijklmnopqrstuvwxyz", "bcadefghijklmnopqrstuvwxzz", true},
		{"aabcc", "ccdee", true},
		{"leetcode", "codeleet", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.str1), func(t *testing.T) {
			require.Equal(t, tc.want, canConvert(tc.str1, tc.str2))
		})
	}
}

func canConvert(str1 string, str2 string) bool {
	// If the second string contains the entire alphabet, then it is not possible
	// to swap str1 to str2 even if a mapping exists. This is because there is no
	// way to swap one character to another without overlap with existing ones.
	var charSeen int
	for i := range str2 {
		charSeen |= 1 << (str2[i] - 'a')
	}
	if charSeen == (1<<26)-1 {
		return str1 == str2
	}

	var mapping [26]byte
	for i := range str1 {
		if mapping[str1[i]-'a'] == 0 {
			mapping[str1[i]-'a'] = str2[i]
		} else {
			if mapping[str1[i]-'a'] != str2[i] {
				return false
			}
		}
	}
	return true
}

package p0953verifyinganaliendict

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isAlienSorted(t *testing.T) {
	for _, tc := range []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz", true},
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, isAlienSorted(tc.words, tc.order))
		})
	}
}

func isAlienSorted(words []string, order string) bool {
	var alienPosition [26]byte
	for i, o := range order {
		alienPosition[byte(o)-'a'] = byte(i)
	}
	for i := range words[1:] {
		for j := range words[i] {
			if j >= len(words[i+1]) {
				// Match except next word is greater -> fail
				return false
			}
			a, b := words[i][j]-'a', words[i+1][j]-'a'
			i, j := alienPosition[a], alienPosition[b]
			if i < j {
				break // done
			} else if i > j {
				return false
			}
		}
	}
	return true
}

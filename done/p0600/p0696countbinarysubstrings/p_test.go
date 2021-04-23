package p0696countbinarysubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBinarySubstrings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"00110011", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countBinarySubstrings(tc.s))
		})
	}
}

func countBinarySubstrings(s string) int {
	var pos int
	n := len(s)
	var res int
	var zeroes, ones int
	for pos < n {
		zeroes = 0
		for pos < n && s[pos] == '0' {
			zeroes++
			if ones > 0 {
				res++
				ones--
			}
			pos++
		}
		ones = 0
		for pos < n && s[pos] == '1' {
			ones++
			if zeroes > 0 {
				res++
				zeroes--
			}
			pos++
		}
	}
	return res
}

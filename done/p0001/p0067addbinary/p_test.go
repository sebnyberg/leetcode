package p0067addbinary

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBinary(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"100", "110010", "110110"},
		// {"11", "1", "100"},
		// {"0", "1", "1"},
		// {"0", "0", "0"},
		// {"1010", "1011", "10101"},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.a, tc.b), func(t *testing.T) {
			require.Equal(t, tc.want, addBinary(tc.a, tc.b))
		})
	}
}

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	na := len(a)
	nb := len(b)
	var carry int
	var resp strings.Builder
	for i := 0; i < len(b); i++ {
		carry += int(a[na-1-i]) - '0' + int(b[nb-1-i]) - '0'
		resp.WriteRune(rune(carry%2 + '0'))
		carry /= 2
	}
	for i := nb; i < na; i++ {
		carry += int(a[na-1-i] - '0')
		resp.WriteRune(rune(carry%2 + '0'))
		carry /= 2
	}
	if carry == 1 {
		resp.WriteRune('1')
	}
	runes := []rune(resp.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

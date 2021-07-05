package p1147longestchunkedpalindromedecomposition

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestDecomposition(t *testing.T) {
	for _, tc := range []struct {
		text string
		want int
	}{
		{"ghiabcdefhelloadamhelloabcdefghi", 7},
		{"merchant", 1},
		{"antaprezatepzapreanta", 11},
		{"aaa", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.text), func(t *testing.T) {
			require.Equal(t, tc.want, longestDecomposition(tc.text))
		})
	}
}

func longestDecomposition(text string) int {
	n := len(text)
	mod := 1_000_000_007
	base := 29
	h1, h2 := 0, 0
	pow := 1
	var res int
	for i := 0; i < n/2; i++ {
		h1 = (h1*base + int(text[i]-'a')) % mod
		h2 = (pow*int(text[n-i-1]-'a') + h2) % mod
		pow = pow * base % mod
		if h1 == h2 {
			res += 2
			h1 = 0
			h2 = 0
			pow = 1
		}
	}
	if n%2 == 1 || h1 != h2 {
		res += 1
	}
	return res
}

package p0541reversestring2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseStr(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, reverseStr(tc.s, tc.k))
		})
	}
}

func reverseStr(s string, k int) string {
	n := len(s)
	bs := []byte(s)
	for i := 0; i < n; i += 2 * k {
		if i+2*k >= n {
			if i+k > n {
				for l, r := i, n-1; l < r; l, r = l+1, r-1 {
					bs[l], bs[r] = bs[r], bs[l]
				}
			} else {
				for l, r := i, i+k-1; l < r; l, r = l+1, r-1 {
					bs[l], bs[r] = bs[r], bs[l]
				}
			}
		} else {
			for l, r := i, i+k-1; l < r; l, r = l+1, r-1 {
				bs[l], bs[r] = bs[r], bs[l]
			}
		}
	}
	return string(bs)
}

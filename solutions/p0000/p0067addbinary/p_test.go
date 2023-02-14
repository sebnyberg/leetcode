package p0067addbinary

import (
	"fmt"
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
	m := max(len(a), len(b))
	res := make([]byte, 0, m)
	var carry byte
	for i := 0; i < m; i++ {
		ai := len(a) - i - 1
		bi := len(b) - i - 1
		var aa, bb byte
		if ai >= 0 {
			aa = byte(a[ai] - '0')
		}
		if bi >= 0 {
			bb = byte(b[bi] - '0')
		}
		v := aa + bb + carry
		carry = v / 2
		res = append(res, v&1+'0')
	}
	if carry > 0 {
		res = append(res, '1')
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

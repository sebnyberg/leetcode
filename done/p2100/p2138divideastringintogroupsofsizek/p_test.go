package p2138divideastringintogroupsofsizek

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divideString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		fill byte
		want []string
	}{
		{"abcdefghi", 3, 'x', []string{"abc", "def", "ghi"}},
		{"abcdefghij", 3, 'x', []string{"abc", "def", "ghi", "jxx"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, divideString(tc.s, tc.k, tc.fill))
		})
	}
}

func divideString(s string, k int, fill byte) []string {
	n := len(s)
	res := make([]string, 0, n/3)
	for i := 0; i < len(s); i += k {
		res = append(res, s[i:min(n, i+k)])
		if w := n - i; w < k {
			res[len(res)-1] += string(bytes.Repeat([]byte{fill}, k-w))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

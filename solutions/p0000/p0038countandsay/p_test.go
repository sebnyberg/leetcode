package p0038countandsay

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_countAndSay(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want string
	}{
		{1, "1"},
		{4, "1211"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, countAndSay(tc.in))
		})
	}
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	res := make([]byte, 0)
	var end int
	for start := 0; start < len(s); start = end {
		for end = start + 1; end < len(s) && s[end] == s[start]; end++ {
		}
		res = append(res, byte(end-start+'0'), s[start])
	}
	return *(*string)(unsafe.Pointer(&res))
}

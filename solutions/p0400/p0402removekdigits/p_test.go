package p0402removekdigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeKdigits(t *testing.T) {
	for _, tc := range []struct {
		num  string
		k    int
		want string
	}{
		{"10", 1, "0"},
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, removeKdigits(tc.num, tc.k))
		})
	}
}

func removeKdigits(num string, k int) string {
	n := len(num)
	want := n - k
	if k == len(num) {
		return "0"
	}

	stack := make([]byte, 0)
	m := len(stack)
	for i := range num {
		ch := num[i]
		for m > 0 && ch < stack[m-1] && k > 0 {
			stack = stack[:m-1]
			k--
			m--
		}
		stack = append(stack, ch)
		m++
	}

	// At this point, the stack contains at least 'want' characters
	// Trim prefix characters
	stack = stack[:want]
	var i int
	for i < len(stack)-1 && stack[i] == '0' {
		i++
	}
	return string(stack[i:])
}

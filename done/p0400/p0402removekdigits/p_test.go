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
	// n := len(num)
	// stack := make([]byte, 0)
	// for i, ch := range num {
	// 	if i < k {
	// 		stack = append(stack, byte(ch-'0'))
	// 		continue
	// 	}
	// 	canPop := n - 1
	// for canPop > 0 && len(stack) > 0 && stack[len(stack)-1] > byte(ch-'0') {
	// 	stack = stack[:len(stack)-1]
	// }
	// }
	return ""
}

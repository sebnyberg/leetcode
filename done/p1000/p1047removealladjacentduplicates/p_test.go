package p1047removealladjacentduplicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removeDuplicates(tc.s))
		})
	}
}

func removeDuplicates(s string) string {
	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

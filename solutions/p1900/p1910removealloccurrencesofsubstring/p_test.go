package p1910removealloccurrencesofsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s    string
		part string
		want string
	}{
		{"daabcbaabcbc", "abc", "dab"},
		{"axxxxyyyyb", "xy", "ab"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removeOccurrences(tc.s, tc.part))
		})
	}
}

func removeOccurrences(s string, part string) string {
	stack := make([]rune, 0)
	for _, ch := range s {
		stack = append(stack, ch)
		if len(stack) >= len(part) {
			if string(stack[len(stack)-len(part):]) == part {
				stack = stack[:len(stack)-len(part)] // Pop
			}
		}
	}
	return string(stack)
}

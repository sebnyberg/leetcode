package p0020validparen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, isValid(tc.in))
		})
	}
}

func isValid(s string) bool {
	opened := make([]rune, 0)
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			opened = append(opened, ch)
		case ')':
			if len(opened) == 0 || opened[len(opened)-1] != '(' {
				return false
			}
			opened = opened[:len(opened)-1]
		case ']':
			if len(opened) == 0 || opened[len(opened)-1] != '[' {
				return false
			}
			opened = opened[:len(opened)-1]
		case '}':
			if len(opened) == 0 || opened[len(opened)-1] != '{' {
				return false
			}
			opened = opened[:len(opened)-1]
		}
	}
	return len(opened) == 0
}

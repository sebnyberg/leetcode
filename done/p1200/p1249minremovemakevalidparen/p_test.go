package p1249minremovemakevalidparen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
		{"(a(b(c)d)", "a(b(c)d)"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minRemoveToMakeValid(tc.s))
		})
	}
}

func minRemoveToMakeValid(s string) string {
	// Create a map of indices to skip
	// Indices to skip are either leading ')'s or trailing '('s
	lparens := make([]int, 0)      // (stack) of left paren indices
	skip := make(map[int]struct{}) // map of indices to skip
	for i, ch := range s {
		switch ch {
		case '(':
			lparens = append(lparens, i) // add to stack
		case ')':
			if len(lparens) == 0 {
				skip[i] = struct{}{} // leading ')'
			} else {
				// Pop '(' from the stack
				lparens = lparens[:len(lparens)-1]
			}
		}
	}

	// Add remaining '(' indices to the map of indices to skip
	for _, pos := range lparens {
		skip[pos] = struct{}{}
	}

	// Create resulting string
	var res strings.Builder
	for i, ch := range s {
		if _, exists := skip[i]; exists {
			continue
		}
		res.WriteRune(ch)
	}
	return res.String()
}

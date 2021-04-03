package plongestvalidparen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestValidParentheses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{")()())", 4},
		{"(()", 2},
		{"", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestValidParentheses(tc.s))
		})
	}
}

type stack []parenPos

func (s stack) len() int { return len(s) }
func (s *stack) pop() parenPos {
	n := s.len()
	it := (*s)[n-1]
	(*s) = (*s)[:n-1]
	return it
}
func (s *stack) push(n parenPos) { (*s) = append((*s), n) }
func (s *stack) peek() parenPos  { return (*s)[len(*s)-1] }

type parenPos struct {
	paren byte
	pos   int
}

func longestValidParentheses(s string) int {
	var parens stack
	var maxLen int
	for i := range s {
		switch s[i] {
		case ')':
			if parens.len() == 0 {
				parens.push(parenPos{')', i})
				continue
			}
			prev := parens.peek()
			if prev.paren == '(' {
				parens.pop()
				if parens.len() == 0 {
					maxLen = max(maxLen, i+1)
				} else {
					maxLen = max(maxLen, i-parens.peek().pos)
				}
				continue
			}
			parens.push(parenPos{')', i})
		case '(':
			parens.push(parenPos{'(', i})
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

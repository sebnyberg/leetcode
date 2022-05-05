package p_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestValidParenthesis(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"(()", 2},
		{"())", 2},
		{")()())", 4},
		{"()(())", 6},
		{"))(())()())", 8},
		{"(()()", 4},
		{")()())()()(", 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, longestValidParentheses(tc.in))
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	opened := make([]int, 0)
	var maxvalid int
	var nopen int
	for i, ch := range s {
		switch ch {
		case '(':
			opened = append(opened, i)
			nopen++
		case ')':
			switch len(opened) {
			case 0:
				nopen = 0
				continue
			case 1:
				maxvalid = max(maxvalid, nopen*2)
			default:
				d := i - opened[len(opened)-2]
				maxvalid = max(maxvalid, d)
			}
			opened = opened[:len(opened)-1]
		}
	}
	return maxvalid
}

//////////////////////////////////////////////////////////
// Stack based version

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

func longestValidParenthesesStack(s string) int {
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

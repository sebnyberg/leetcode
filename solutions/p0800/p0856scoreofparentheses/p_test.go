package p2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_scoreOfParentheses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"()()", 2},
		{"(())", 2},
		{"()", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, scoreOfParentheses(tc.s))
		})
	}
}

func scoreOfParentheses(s string) int {
	const lparen = -2
	stack := []int{}
	var n int
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, lparen)
			n++
			continue
		}
		if stack[n-1] == lparen {
			stack[n-1] = 1
		} else {
			stack[n-2] = stack[n-1] * 2
			stack = stack[:n-1]
			n--
		}
		// Combine with any prior values
		for len(stack) > 1 && stack[n-2] != lparen {
			stack[n-2] += stack[n-1]
			stack = stack[:n-1]
			n--
		}
	}
	return stack[0]
}

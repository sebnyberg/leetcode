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

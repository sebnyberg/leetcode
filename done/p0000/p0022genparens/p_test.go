package p0022genparens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateParenthesis(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want []string
	}{
		// {3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, generateParenthesis(tc.in))
		})
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	genParenthesis(&res, "", n, n)
	return res
}

func genParenthesis(res *[]string, s string, toOpen, toClose int) {
	if toClose == 0 {
		*res = append(*res, s)
		return
	}
	if toOpen > 0 {
		genParenthesis(res, s+"(", toOpen-1, toClose)
	}
	if toOpen < toClose {
		genParenthesis(res, s+")", toOpen, toClose-1)
	}
}

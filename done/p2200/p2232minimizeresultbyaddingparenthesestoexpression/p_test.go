package p2232minimizeresultbyaddingparenthesestoexpression

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizeResult(t *testing.T) {
	for _, tc := range []struct {
		expression string
		want       string
	}{
		{"12+34", "1(2+3)4"},
		{"247+38", "2(47+38)"},
		{"999+999", "(999+999)"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeResult(tc.expression))
		})
	}
}

func minimizeResult(expression string) string {
	parts := strings.Split(expression, "+")
	l, r := parts[0], parts[1]
	res := "(" + expression + ")"
	minVal := parseInt(l) + parseInt(r)
	for i := 0; i < len(l); i++ {
		for j := len(r); j > 0; j-- {
			fac := 1
			fac *= max(1, parseInt(l[:i]))
			fac *= max(1, parseInt(r[j:]))
			term := parseInt(l[i:]) + parseInt(r[:j])
			if fac*term < minVal {
				minVal = fac * term
				res = l[:i] + "(" + l[i:] + "+" + r[:j] + ")" + r[j:]
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	var x int
	for _, ch := range s {
		x *= 10
		x += int(ch - '0')
	}
	return x
}

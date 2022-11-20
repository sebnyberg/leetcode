package p0224basiccalculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculate(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"(6)-(8)-(7)+(1+(6))", -2},
		{"1-(     -2)", 3},
		{" 2-1 + 2 ", 3},
		{"1 + 1", 2},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"- (3 + (4 + 5))", -12},
		{"- (1)", -1},
		{"-2+ 1", -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, calculate(tc.s))
		})
	}
}

func calculate(s string) int {
	// val and sign contains the value and last sign
	// on '(', we push to val and sign
	// on ')', we pop and use sign of prior expression
	val := []int{0}
	sign := []int{1}
	var j int // stack position (expression depth)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '-':
			sign[j] = -1
		case '+', ' ':
			continue
		case '(':
			val = append(val, 0)
			sign = append(sign, 1)
			j++
		case ')':
			val[j-1] = val[j-1] + sign[j-1]*val[j]
			val = val[:j]
			sign = sign[:j]
			j--
			sign[j] = 1
		default: // number
			var x int
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				x = x*10 + int(s[i]-'0')
				i++
			}
			i-- // loop will increment i once more
			val[j] = val[j] + sign[j]*x
			sign[j] = 1
		}
	}
	return val[0]
}

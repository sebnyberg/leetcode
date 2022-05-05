package p0439ternaryexpressionparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseTernary(t *testing.T) {
	for _, tc := range []struct {
		expression string
		want       string
	}{
		{"T?2:3", "2"},
		{"F?1:T?4:5", "4"},
		{"T?T?F:4:3", "F"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
			require.Equal(t, tc.want, parseTernary(tc.expression))
		})
	}
}

func parseTernary(expression string) string {
	var vals, syms []byte
	expression += ":"
	for i := 0; i < len(expression); i += 2 {
		vals = append(vals, expression[i])
		syms = append(syms, expression[i+1])
		j := len(vals) - 3
		for len(vals) >= 3 &&
			syms[j] == '?' && syms[j+1] == ':' && syms[j+2] == ':' {
			if vals[j] == 'T' {
				vals[j] = vals[j+1]
			} else {
				vals[j] = vals[j+2]
			}
			syms[j] = ':'
			vals = vals[:j+1]
			syms = syms[:j+1]
			j = len(vals) - 3
		}
	}
	return string(vals[0])
}

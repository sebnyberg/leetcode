package p0640solvetheequation

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_solveEquation(t *testing.T) {
	for i, tc := range []struct {
		equation string
		want     string
	}{
		{"x+5-3+x=6+x-2", "x=2"},
		{"x=x", "Infinite solutions"},
		{"2x=x", "x=0"},
		{"x=x+2", "No solution"},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, solveEquation(tc.equation))
		})
	}
}

func solveEquation(equation string) string {
	// ax+b
	gather := func(s string) (b int, a int) {
		// + or - or nothing
		// numeric until end or non-numeric
		// x or not x
		var i int
		sign := 1
		for i < len(s) {
			if s[i] == '+' {
				i++
				sign = 1
				continue
			}
			if s[i] == '-' {
				i++
				sign = -1
				continue
			}
			var res int
			var fdsafsd bool
			for i < len(s) && unicode.IsNumber(rune(s[i])) {
				fdsafsd = true
				res *= 10
				res += int(s[i] - '0')
				i++
			}
			if res == 0 && !fdsafsd {
				res = 1
			}
			if i < len(s) && s[i] == 'x' {
				a += sign * res
				i++
			} else {
				b += sign * res
			}
		}
		return a, b
	}
	parts := strings.Split(equation, "=")
	a1, b1 := gather(parts[0])
	a2, b2 := gather(parts[1])
	a := a1 - a2
	b := b2 - b1
	if a == 0 && b == 0 {
		return "Infinite solutions"
	}
	if a == 0 {
		return "No solution"
	}
	if a < 0 {
		a = -a
		b = -b
	}
	res := fmt.Sprintf("x=%v", b/a)
	return res
}

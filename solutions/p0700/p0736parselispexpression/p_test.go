package p0736parselispexpression

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_evaluate(t *testing.T) {
	for _, tc := range []struct {
		expression string
		want       int
	}{
		{"(let x 2 (add (let x 3 (let x 4 x)) x))", 6},
		// {"(let x 2 (mult x (let x 3 y 4 (add x y))))", 14},
		// {"(let x 3 x 2 x)", 2},
		// {"(let x 1 y 2 x (add x y) (add x y))", 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
			require.Equal(t, tc.want, evaluate(tc.expression))
		})
	}
}

// A symbol either contains a set of symbols, or a value
type symbol struct {
	contents string
	symbols  []*symbol
}

func evaluate(expression string) int {
	// Parse symbols
	// Each symbol can either contain a value, or a list of symbols, never both.
	// Symbol contents are either a variable name, or a value
	root := symbol{}
	stack := []*symbol{&root}
	for i := 1; i < len(expression)-1; {
		ch := expression[i]
		switch ch {
		case ' ':
			i++
		case '(':
			var s symbol
			stack[len(stack)-1].symbols = append(stack[len(stack)-1].symbols, &s)
			stack = append(stack, &s)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			// Read symbol
			var s symbol
			j := strings.IndexAny(expression[i:], "() ")
			s.contents = expression[i : i+j]
			i += j
			stack[len(stack)-1].symbols = append(stack[len(stack)-1].symbols, &s)
		}
	}

	mustParseInt := func(s string) int {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return x
	}

	var eval func(map[string]int, *symbol) int
	eval = func(m map[string]int, s *symbol) int {
		mm := make(map[string]int, len(m))
		for k, v := range m {
			mm[k] = v
		}
		m = mm
		if len(s.symbols) == 0 {
			if v, exists := m[s.contents]; exists {
				return v
			}
			return mustParseInt(s.contents)
		}
		switch s.symbols[0].contents {
		case "add":
			return eval(m, s.symbols[1]) + eval(m, s.symbols[2])
		case "mult":
			return eval(m, s.symbols[1]) * eval(m, s.symbols[2])
		case "let":
			for i := 1; i+2 <= len(s.symbols); i += 2 {
				m[s.symbols[i].contents] = eval(m, s.symbols[i+1])
			}
			return eval(m, s.symbols[len(s.symbols)-1])
		}

		panic("end of symbol switch")
	}

	m := make(map[string]int)
	res := eval(m, &root)

	return res
}

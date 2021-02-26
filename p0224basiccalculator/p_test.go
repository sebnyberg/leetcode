package p0224basiccalculator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculate(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"(6)-(8)-(7)+(1+(6))", -2},
		{" 2-1 + 2 ", 3},
		{"1 + 1", 2},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"- (3 + (4 + 5))", -12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, calculate(tc.s))
		})
	}
}

func calculate(s string) int {
	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	c := calculator{
		input: s,
	}
	return c.addExpr()
}

// Grammar:
// addExpr := subExpr { ("+") addExpr }
// subExpr := term { ("-") subExpr }
// term    := factor { ("*" | "/") term }
// factor  := "-" num | "-" "(" expr ")"
type calculator struct {
	input string
	pos   int
}

var eof = rune(0)

func (c *calculator) addExpr() int {
	lhs := c.subExpr()
	switch c.next() {
	case eof:
		return lhs
	case '+':
		return lhs + c.addExpr()
	default: //  end of expression
		c.pos--
		return lhs
	}
}

func (c *calculator) subExpr() int {
	lhs := c.term()
	switch c.next() {
	case eof:
		return lhs
	case '-':
		return lhs - c.subExpr()
	default: // end of expression
		c.pos--
		return lhs
	}
}

func (c *calculator) term() int {
	lhs := c.factor()
	switch c.next() {
	case eof:
		return lhs
	case '*':
		return lhs * c.term()
	case '/':
		return lhs / c.term()
	default: // bubble up when the operator was '+' or '-'
		c.pos--
		return lhs
	}
}

func (c *calculator) factor() int {
	sign := 1
	if c.next() == '-' {
		sign = -1
	} else {
		c.pos--
	}

	switch r := c.next(); {
	case r == '(':
		inner := c.addExpr()
		c.next() // parse end parenthesis
		return sign * inner
	case strings.ContainsRune("0123456789", r):
		start := c.pos - 1
		for strings.ContainsRune("0123456789", c.next()) {
		}
		c.pos-- // back off last rune
		n, err := strconv.Atoi(c.input[start:c.pos])
		if err != nil {
			log.Fatalln(err)
		}
		return sign * n
	default:
		log.Fatalln("invalid token", r)
		return 0
	}
}

func (p *calculator) next() rune {
	if p.pos >= len(p.input) {
		p.pos++
		return eof
	}
	r := p.input[p.pos]
	p.pos++
	return rune(r)
}

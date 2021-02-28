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
		{"- (1)", -1},
		{"-2+ 1", -1},
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
		pos:   len(s) - 1,
	}
	return c.expr()
}

// Grammar: (now RL)
// expr    := expr { ("-" | "+") term }
// term    := term { ("*" | "/") factor }
// factor  := "-" num | "-" "(" expr ")"

type calculator struct {
	input string
	pos   int
}

var eof = rune(0)

func (c *calculator) expr() int {
	rhs := c.term()
	switch c.next() {
	case eof:
		return rhs
	case '+':
		return c.expr() + rhs
	case '-':
		return c.expr() - rhs
	default: //  end of expression
		c.backup()
		return rhs
	}
}

func (c *calculator) term() int {
	rhs := c.factor()
	switch c.next() {
	case eof:
		return rhs
	case '*':
		return c.term() * rhs
	case '/':
		return c.term() / rhs
	default: // bubble up when the operator was '+' or '-'
		c.backup()
		return rhs
	}
}

func (c *calculator) factor() int {

	switch r := c.next(); {
	case r == ')':
		inner := c.expr()
		c.next() // parse end parenthesis
		sign := 1
		if c.next() == '-' {
			if r := c.next(); r == '-' || r == eof {
				sign = -1
			} else {
				c.backup()
				c.backup()
			}
		} else {
			c.backup()
		}
		return sign * inner
	case strings.ContainsRune("0123456789", r):
		start := c.pos + 2
		for strings.ContainsRune("0123456789", c.next()) {
		}
		c.backup()
		n, err := strconv.Atoi(c.input[c.pos+1 : start])
		if err != nil {
			log.Fatalln(err)
		}
		sign := 1
		if c.next() == '-' {
			if r := c.next(); r == '-' || r == eof {
				sign = -1
			} else {
				c.backup()
				c.backup()
			}
		} else {
			c.backup()
		}
		return sign * n
	default:
		log.Fatalln("invalid token", r)
		return 0
	}
}

func (p *calculator) next() rune {
	if p.pos < 0 {
		p.pos--
		return eof
	}
	r := p.input[p.pos]
	p.pos--
	return rune(r)
}

func (p *calculator) backup() {
	p.pos++
}

package p0856scoreofparen

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_scoreOfParentheses(t *testing.T) {
	for _, tc := range []struct {
		S    string
		want int
	}{
		{"(())()", 3},
		{"(()(()))", 6},
		{"(())", 2},
		{"()", 1},
		{"()()", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.S), func(t *testing.T) {
			require.Equal(t, tc.want, scoreOfParentheses(tc.S))
		})
	}
}

func scoreOfParentheses(S string) (res int) {
	p := parser{s: S}
	return p.scoreParen()
}

type parser struct {
	s   string
	pos int
}

func (p *parser) scoreParen() int {
	if p.next() == eof {
		return 0
	}

	// p.next() can be either '(' or ')'
	switch p.peek() {
	case ')':
		p.next()
		if p.peek() == '(' {
			return 1 + p.scoreParen()
		} else {
			return 1
		}
	case '(':
		res := 2 * p.scoreParen()
		p.next()
		if p.peek() == '(' {
			return res + p.scoreParen()
		} else {
			return res
		}
	}

	panic("fail")
}

var eof = rune(0)

func (p *parser) next() rune {
	if p.pos >= len(p.s) {
		return eof
	}
	// No need to worry about width - all characters are ASCII
	r, w := utf8.DecodeRuneInString(p.s[p.pos:])
	p.pos += w
	return r
}

func (p *parser) peek() rune {
	if p.pos >= len(p.s) {
		return eof
	}
	r, _ := utf8.DecodeRuneInString(p.s[p.pos:])
	return r
}

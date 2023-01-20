package p1106parsingabooleanexpression

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseBoolExpr(t *testing.T) {
	for i, tc := range []struct {
		expression string
		want       bool
	}{
		{"&(|(f))", false},
		{"|(f,f,f,t)", true},
		{"!(&(f,t))", true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, parseBoolExpr(tc.expression))
		})
	}
}

func parseBoolExpr(expression string) bool {
	var p parser
	p.s = expression
	prog := p.parseExpr()
	res := prog.eval()
	return res
}

const (
	opNot = "!"
	opAnd = "&"
	opOr  = "|"
)

type expr struct {
	val      bool
	subExprs []*expr
	op       string
}

func (e *expr) eval() bool {
	if e.op == "" {
		return e.val
	}
	if len(e.subExprs) == 0 {
		panic("no subexprs")
	}
	v := e.subExprs[0].eval()
	switch e.op {
	case opNot:
		v = !v
	case opAnd:
		for _, s := range e.subExprs {
			v = v && s.eval()
		}
	case opOr:
		for _, s := range e.subExprs {
			v = v || s.eval()
		}
	default:
		log.Fatalf("invalid op %s", e.op)
	}
	return v
}

type parser struct {
	s string
	i int
}

func (p *parser) parseExpr() *expr {
	var e expr
	if p.i >= len(p.s) {
		return nil
	}
	switch p.s[p.i] {
	case 'f':
		e.val = false
	case 't':
		e.val = true
	case '|':
		p.i++
		for p.s[p.i] != ')' {
			p.i++
			e.subExprs = append(e.subExprs, p.parseExpr())
		}
		e.op = opOr
	case '&':
		p.i++
		e.op = opAnd
		for p.s[p.i] != ')' {
			p.i++
			e.subExprs = append(e.subExprs, p.parseExpr())
		}
	case '!':
		p.i++
		e.op = opNot
		for p.s[p.i] != ')' {
			p.i++
			e.subExprs = append(e.subExprs, p.parseExpr())
		}
	default:
		log.Fatalln(p.s[p.i:])
	}
	p.i++

	return &e
}

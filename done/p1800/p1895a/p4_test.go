package p1895a

import (
	"io"
	"log"
	"os"
	"testing"
	"unsafe"
)

func Test_minOperationsToFlip(t *testing.T) {
	// for _, tc := range []struct {
	// 	expression string
	// 	want       int
	// }{
	// 	{"(0&0)&(0&0&0)", 3},
	// 	{"1&(0|1)", 1},
	// 	{"(0|(1|0&1))", 1},
	// } {
	// 	t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
	// 		require.Equal(t, tc.want, minOperationsToFlip(tc.expression))
	// 	})
	// }
	t.Run("big", func(t *testing.T) {
		f, _ := os.Open("testdata/input")
		s, _ := io.ReadAll(f)
		ss := string(s)
		res := minOperationsToFlip(ss)
		_ = res
	})
}

func minOperationsToFlip(expression string) int {
	// Evaluate current value of the expression
	bs := *(*[]byte)(unsafe.Pointer(&expression))
	parsedExpr := parseExpr(bs)
	var want bool
	if !parsedExpr.eval() {
		want = true
	}
	res := minOpToFlip(parsedExpr, want)
	return res
}

func minOpToFlip(e expr, want bool) int {
	if e.eval() == want {
		return 0
	}
	if e.getType() == typUnary {
		return 1
	}
	var res int
	binExpr := e.(*binaryExpr)
	if !want {
		// Expression evaluated to true, although it should be false.
		// If the expression has an OR, it is always worth it to change the OR to
		// to an AND. Consider the three cases where the result is true:
		// 0|1, 1|0, 1|1. Changing 0|1 and 1|0 to false requires either changing
		// the operand (| -> &) or making the true operand false (1 -> 0). In any
		// case where the operand is not just a single number, it is more expensive
		// to change the sub-expression's value rather than the operator.
		// Thus, when the expression is true, always change from OR => AND.
		if binExpr.op == binOpOr {
			res++
		}
		// The operand is guaranteed to be AND. If both are true, turn the
		// cheapest alternative into false.
		if binExpr.lhs.eval() && binExpr.rhs.eval() {
			res += min(minOpToFlip(binExpr.lhs, want), minOpToFlip(binExpr.rhs, want))
		}
	} else {
		// When the expression is false, always change from AND => OR.
		if binExpr.op == binOpAnd {
			res++
		}
		// Now the operand is guaranteed to be OR. If both are false, turn the
		// cheapest alternative into true.
		if !binExpr.lhs.eval() && !binExpr.rhs.eval() {
			res += min(minOpToFlip(binExpr.lhs, want), minOpToFlip(binExpr.rhs, want))
		}
	}
	return res
}

func parseExpr(exprBytes []byte) expr {
	p := parser{expr: exprBytes, n: len(exprBytes), pos: 0}
	return p.parseExpr()
}

type parser struct {
	expr []byte
	n    int
	pos  int
}

// acceptExpr parses an expression. An expression contains a left-hand side and
// zero or many right-hand sides.
func (p *parser) parseExpr() expr {
	var lhs expr

	// Parse LHS
	if p.peek() == '(' {
		p.next()
		lhs = p.parseExpr()
		p.next() // ')'
	} else {
		lhs = unaryExpr(p.next() == '1')
	}

	// Reduce LHS and RHS until end of parenthesis expr or EOF
	for p.peek() != ')' && p.peek() != eof {
		op := binOp(p.next())
		var rhs expr
		nextToken := p.next()
		if nextToken == '(' {
			rhs = p.parseExpr()
			p.next() // ')'
		} else {
			rhs = unaryExpr(nextToken == '1')
		}
		lhs = &binaryExpr{
			lhs: lhs,
			op:  op,
			rhs: rhs,
		}
	}
	return lhs
}

const eof = byte(0)

func (p *parser) next() byte {
	if p.pos >= p.n {
		return eof
	}
	res := p.expr[p.pos]
	p.pos++
	return res
}

func (p *parser) peek() byte {
	if p.pos >= p.n {
		return eof
	}
	return p.expr[p.pos]
}

type typ int

const (
	typBinary typ = 0
	typUnary  typ = 1
)

type expr interface {
	getType() typ
	eval() bool
}

type binOp byte

const (
	binOpAnd binOp = '&'
	binOpOr  binOp = '|'
)

type binaryExpr struct {
	lhs     expr
	rhs     expr
	op      binOp
	val     bool
	didEval bool
}

func (e *binaryExpr) eval() bool {
	if !e.didEval {
		switch e.op {
		case binOpAnd:
			e.val = e.lhs.eval() && e.rhs.eval()
		case binOpOr:
			e.val = e.lhs.eval() || e.rhs.eval()
		default:
			log.Fatalln("invalid op", e.op)
		}
	}
	e.didEval = true
	return e.val
}

func (e binaryExpr) getType() typ { return typBinary }

type unaryExpr bool

func (e unaryExpr) eval() bool   { return bool(e) }
func (e unaryExpr) getType() typ { return typUnary }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

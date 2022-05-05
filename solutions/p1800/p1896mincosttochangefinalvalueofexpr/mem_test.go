package p1896mincosttochangefinalvalueofexpr

// memory optimized

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_minOperationsToFlip(t *testing.T) {
// 	for _, tc := range []struct {
// 		expression string
// 		want       int
// 	}{
// 		{"(0&0)&(0&0&0)", 3},
// 		{"1&(0|1)", 1},
// 		{"(0|(1|0&1))", 1},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
// 			require.Equal(t, tc.want, minOperationsToFlip(tc.expression))
// 		})
// 	}
// 	for i, input := range []string{
// 		"testdata/input",
// 		"testdata/input2",
// 	} {
// 		t.Run(fmt.Sprintf("big_%v", i), func(t *testing.T) {
// 			f, _ := os.Open(input)
// 			s, _ := io.ReadAll(f)

// 			ss := string(s)
// 			res := minOperationsToFlip(ss)
// 			_ = res
// 		})
// 	}
// }

// var res int

// func Benchmark_minOperationsToFlip(b *testing.B) {
// 	f, _ := os.Open("testdata/input")
// 	s, _ := io.ReadAll(f)

// 	ss := string(s)
// 	for i := 0; i < b.N; i++ {
// 		res = minOperationsToFlip(ss)
// 	}
// }

// func minOperationsToFlip(expression string) int {
// 	bs := []byte(expression)
// 	tree := parseExpr(bs)
// 	var want bool
// 	if !tree.eval(tree.rootIdx, 0, false) {
// 		want = true
// 	}
// 	res := minOpToFlip(*tree, tree.rootIdx, want)
// 	return int(res)
// }

// func minOpToFlip(t exprTree, idx uint16, want bool) uint16 {
// 	if t.eval(idx, 0, false) == want {
// 		return 0
// 	}
// 	e := t.exprs[idx]
// 	if e.op == opNop {
// 		if e.lhsIdx == 0 {
// 			return 1
// 		}
// 		return minOpToFlip(t, e.lhsIdx, want)
// 	}

// 	var res uint16
// 	if !want {
// 		// Expression evaluated to true, although it should be false.
// 		// If the expression has an OR, it is always worth it to change the OR to
// 		// to an AND. Consider the three cases where the result is true:
// 		// 0|1, 1|0, 1|1. Changing 0|1 and 1|0 to false requires either changing
// 		// the operand (| -> &) or making the true operand false (1 -> 0). In any
// 		// case where the operand is not just a single number, it is more expensive
// 		// to change the sub-expression's value rather than the operator.
// 		// Thus, when the expression is true, always change from OR => AND.
// 		if e.op == opOr {
// 			res++
// 		}
// 		// The operand is guaranteed to be AND. If both are true, turn the
// 		// cheapest alternative into false.
// 		if t.eval(e.lhsIdx, idx, true) && t.eval(e.rhsIdx, idx, false) {
// 			if e.lhsIdx == 0 {
// 				res++
// 			} else if e.rhsIdx == 0 {
// 				res++
// 			} else {
// 				res += min(
// 					minOpToFlip(t, e.lhsIdx, want),
// 					minOpToFlip(t, e.rhsIdx, want),
// 				)
// 			}
// 		}
// 	} else {
// 		// When the expression is false, always change from AND => OR.
// 		if e.op == opAnd {
// 			res++
// 		}
// 		// Now the operand is guaranteed to be OR. If both are false, turn the
// 		// cheapest alternative into true.
// 		if !t.eval(e.lhsIdx, idx, true) && !t.eval(e.rhsIdx, idx, false) {
// 			if e.lhsIdx == 0 {
// 				res++
// 			} else if e.rhsIdx == 0 {
// 				res++
// 			} else {
// 				res += min(
// 					minOpToFlip(t, e.lhsIdx, want),
// 					minOpToFlip(t, e.rhsIdx, want),
// 				)
// 			}
// 		}
// 	}
// 	return res
// }

// func min(a, b uint16) uint16 {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func parseExpr(exprBytes []byte) *exprTree {
// 	p := parser{
// 		expr: exprBytes,
// 		n:    len(exprBytes),
// 		pos:  0,
// 		tree: &exprTree{
// 			exprs:   make([]expr, 1, 1000),
// 			n:       1,
// 			rootIdx: 1,
// 		},
// 	}
// 	p.tree.rootIdx = p.parseExpr()
// 	p.tree.evaled = make([]bool, p.tree.n)
// 	p.tree.values = make([]bool, p.tree.n)
// 	return p.tree
// }

// type parser struct {
// 	expr []byte
// 	n    int
// 	pos  int
// 	tree *exprTree
// }

// type exprTree struct {
// 	exprs   []expr
// 	evaled  []bool
// 	values  []bool
// 	n       uint16
// 	rootIdx uint16
// }

// // acceptExpr parses an expression. An expression contains a left-hand side and
// // zero or many right-hand sides.
// func (p *parser) parseExpr() uint16 {
// 	var e expr
// 	idx := p.tree.n
// 	p.tree.exprs = append(p.tree.exprs, e)
// 	p.tree.n++

// 	// Parse LHS
// 	if p.peek() == '(' {
// 		p.next()
// 		p.tree.exprs[idx].lhsIdx = p.tree.n
// 		p.tree.exprs[idx].lhsIdx = p.parseExpr()
// 		p.next() // ')'
// 	} else {
// 		p.tree.exprs[idx].lhsVal = p.next() == '1'
// 	}

// 	// Reduce LHS and RHS until end of parenthesis expr or EOF
// 	for p.peek() != ')' && p.peek() != eof {
// 		if p.next() == '&' {
// 			p.tree.exprs[idx].op = opAnd
// 		} else {
// 			p.tree.exprs[idx].op = opOr
// 		}
// 		nextToken := p.next()
// 		if nextToken == '(' {
// 			p.tree.exprs[idx].rhsIdx = p.parseExpr()
// 			p.next() // ')'
// 		} else {
// 			p.tree.exprs[idx].rhsVal = nextToken == '1'
// 		}
// 		if p.peek() == ')' || p.peek() == eof {
// 			break
// 		}
// 		p.tree.exprs = append(p.tree.exprs, expr{
// 			lhsIdx: idx,
// 		})
// 		idx = p.tree.n
// 		p.tree.n++
// 	}
// 	return idx
// }

// const eof = byte(0)

// func (p *parser) next() byte {
// 	if p.pos >= p.n {
// 		return eof
// 	}
// 	res := p.expr[p.pos]
// 	p.pos++
// 	return res
// }

// func (p *parser) peek() byte {
// 	if p.pos >= p.n {
// 		return eof
// 	}
// 	return p.expr[p.pos]
// }

// type binOp byte

// const (
// 	opNop binOp = iota
// 	opAnd
// 	opOr
// )

// type expr struct {
// 	lhsIdx, rhsIdx uint16
// 	op             binOp
// 	lhsVal, rhsVal bool
// }

// func (t *exprTree) eval(i uint16, parent uint16, left bool) bool {
// 	if i == 0 {
// 		if left {
// 			return t.exprs[parent].lhsVal
// 		} else {
// 			return t.exprs[parent].rhsVal
// 		}
// 	}
// 	e := t.exprs[i]
// 	if !t.evaled[i] {
// 		t.values[i] = t.eval(e.lhsIdx, i, true)
// 		if e.op != opNop {
// 			if t.exprs[i].op == opAnd {
// 				t.values[i] = t.values[i] && t.eval(e.rhsIdx, i, false)
// 			} else if t.exprs[i].op == opOr {
// 				t.values[i] = t.values[i] || t.eval(e.rhsIdx, i, false)
// 			}
// 		}
// 	}
// 	t.evaled[i] = true
// 	return t.values[i]
// }

package p0301removeinvalidparenthesis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeInvalidParentheses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{")(f", []string{"f"}},
		{"((", []string{""}},
		{"()())()", []string{"(())()", "()()()"}},
		{"(a)())()", []string{"(a())()", "(a)()()"}},
		{")(", []string{""}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, removeInvalidParentheses(tc.s))
		})
	}
}

func removeInvalidParentheses(s string) []string {
	var lparens stack
	var rparens stack
	for i, ch := range s {
		switch ch {
		case '(':
			lparens.push(i)
		case ')':
			if len(lparens) > 0 {
				lparens.pop()
			} else {
				rparens.push(i)
			}
		}
	}

	pf := &ParenFixer{
		stack{},
		stack{},
		s,
		len(s),
		make([]byte, len(s)),
		map[string]struct{}{},
	}
	pf.helper(len(lparens), len(rparens), 0, 0)
	res := make([]string, 0)
	for s := range pf.res {
		res = append(res, s)
	}
	return res
}

type ParenFixer struct {
	lparens stack
	rparens stack
	s       string
	n       int
	buf     []byte
	res     map[string]struct{}
}

// We may skip up to lskip lparens, and rskip rparens
// Once pos is == pf.n, check if current stacks are empty and lskip/rskip
// == 0, if so, then we have found a match
func (pf *ParenFixer) helper(lskip, rskip, origPos, resultPos int) {
	if origPos == pf.n {
		if len(pf.lparens) == 0 && len(pf.rparens) == 0 && lskip == 0 && rskip == 0 {
			k := string(pf.buf[:resultPos])
			pf.res[k] = struct{}{}
		}
		return
	}

	switch pf.s[origPos] {
	case '(':
		// try to not remove the lparen, i.e. take it into account
		pf.lparens.push(resultPos)
		pf.buf[resultPos] = '('
		pf.helper(lskip, rskip, origPos+1, resultPos+1)
		pf.lparens.pop()

		if lskip <= 0 {
			break
		}
		// try to remove (skip) the lparen, i.e. do not take it into account
		pf.helper(lskip-1, rskip, origPos+1, resultPos)

	case ')':
		// try to not remove the rparen, i.e. take it into account
		if len(pf.lparens) > 0 {
			popped := pf.lparens.pop()
			pf.buf[resultPos] = ')'
			pf.helper(lskip, rskip, origPos+1, resultPos+1)
			pf.lparens.push(popped)
		} else {
			pf.rparens.push(resultPos)
			pf.buf[resultPos] = ')'
			pf.helper(lskip, rskip, origPos+1, resultPos+1)
			pf.rparens.pop()
		}

		if rskip <= 0 { // cannot skip any rparens
			break
		}
		// try to remove (skip) the rparen, i.e. do not take it into account
		pf.helper(lskip, rskip-1, origPos+1, resultPos)

	default:
		pf.buf[resultPos] = pf.s[origPos]
		pf.helper(lskip, rskip, origPos+1, resultPos+1)
	}
}

type stack []int

func (s *stack) pop() int {
	n := len(*s)
	it := (*s)[n-1]
	*s = (*s)[:n-1]
	return it
}

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

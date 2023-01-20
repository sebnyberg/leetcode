package p1096braceexpansionii

import (
	"fmt"
	"sort"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_braceExpansionII(t *testing.T) {
	for i, tc := range []struct {
		expression string
		want       []string
	}{
		{"{a,b}{c,{d,e}}", []string{"ac", "ad", "ae", "bc", "bd", "be"}},
		{"{{a,z},a{b,c},{ab,z}}", []string{"a", "ab", "ac", "z"}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, braceExpansionII(tc.expression))
		})
	}
}

func braceExpansionII(expression string) []string {
	var p parser
	p.s = expression
	prog := p.parseProgram()
	m := prog.eval()
	var res []string
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

type program struct {
	cross *crossJoin
}

func (p *program) eval() map[string]struct{} {
	return p.cross.eval()
}

type crossJoin struct {
	sets []*set
}

func (c *crossJoin) eval() map[string]struct{} {
	if len(c.sets) == 0 {
		return map[string]struct{}{}
	}
	curr := c.sets[0].eval()
	next := make(map[string]struct{})
	for _, s := range c.sets[1:] {
		for k := range next {
			delete(next, k)
		}
		for k := range curr {
			for k2 := range s.eval() {
				next[k+k2] = struct{}{}
			}
		}
		curr, next = next, curr
	}
	return curr
}

type set struct {
	val   string
	inner *union
}

func (s *set) eval() map[string]struct{} {
	if s.inner != nil {
		return s.inner.eval()
	}
	return map[string]struct{}{s.val: {}}
}

type union struct {
	crosses []*crossJoin
}

func (u *union) eval() map[string]struct{} {
	res := make(map[string]struct{})
	for _, c := range u.crosses {
		for k := range c.eval() {
			res[k] = struct{}{}
		}
	}
	return res
}

type parser struct {
	s string
	i int
}

func (p *parser) parseProgram() *program {
	return &program{
		cross: p.parseCross(),
	}
}

func (p *parser) parseCross() *crossJoin {
	var res crossJoin
	for {
		s := p.parseSet()
		if s == nil {
			break
		}
		res.sets = append(res.sets, s)
	}
	return &res
}

func (p *parser) parseSet() *set {
	var res set
	if p.i == len(p.s) {
		return nil
	}
	if p.s[p.i] == '{' {
		p.i++ // left {
		res.inner = p.parseUnion()
		p.i++ // right }
	} else if unicode.IsLetter(rune(p.s[p.i])) {
		res.val = string(p.s[p.i]) // single character
		p.i++
	} else {
		return nil
	}
	return &res
}

func (p *parser) parseUnion() *union {
	var res union
	res.crosses = append(res.crosses, p.parseCross())
	for p.i < len(p.s) && p.s[p.i] == ',' {
		p.i++ // ','
		res.crosses = append(res.crosses, p.parseCross())
	}
	return &res
}

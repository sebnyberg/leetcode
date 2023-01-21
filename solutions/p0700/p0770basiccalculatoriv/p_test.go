package p0770basiccalculatoriv

import (
	"sort"
	"strings"
)

type facstr string

func (f facstr) mul(x facstr) facstr {
	a := strings.Split(string(f), "*")
	a = append(a, string(x))
	sort.Strings(a)
	return facstr(strings.Join(a, "*"))
}

// Value represents a set of
type value struct {
	vals map[facstr]int
}

func (v value) mul(b value) *value {
	var res value
	res.vals = make(map[facstr]int)
	for v, coef := range v.vals {
		for v2, coef2 := range b.vals {
			c := coef * coef2
			f := v.mul(v2)
			res.vals[f] += c
		}
	}
	return &res
}

type term struct {
	val   *value
	inner *expr
}

func (t *term) eval() *value {
	if t.inner == nil {
		return t.val
	}
	return t.inner.eval()
}

type expr struct {
	facs []*factor
}

func (e *expr) eval() *value {
	var v value
}

type factor struct {
	terms []*term
	op    byte
}

func (f *factor) eval() *value {
	var v value
	first := f.terms[0].eval()

}

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {

}

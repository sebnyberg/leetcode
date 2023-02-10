package p0770basiccalculatoriv

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_basicCalculatorIV(t *testing.T) {
	for i, tc := range []struct {
		expression string
		evalvars   []string
		evalints   []int
		want       []string
	}{
		{"a + b", []string{"a", "b"}, []int{10, -7}, []string{"3"}},
		{"a * b * c + b * a * c * 4", []string{}, []int{}, []string{"5*a*b*c"}},
		{"(e + 8) * (e - 8)", []string{}, []int{}, []string{"1*e*e", "-64"}},
		{"e - 8 + temperature - pressure", []string{"e", "temperature"}, []int{1, 12}, []string{"-1*pressure", "5"}},
		{"e + 8 - a + 5", []string{"e"}, []int{1}, []string{"-1*a", "14"}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, basicCalculatorIV(tc.expression, tc.evalvars, tc.evalints))
		})
	}
}

const eps = "" // represents a nil-variable, used for constant values

func sortstr(s facstr) facstr {
	a := strings.Split(string(s), "*")
	sort.Strings(a)
	return facstr(strings.Join(a, "*"))
}

type facstr string

func (f facstr) mul(x facstr) facstr {
	if f == eps {
		return x
	}
	if x == eps {
		return f
	}
	a := strings.Split(string(f), "*")
	a = append(a, string(x))
	sort.Strings(a)
	return facstr(strings.Join(a, "*"))
}

type value map[facstr]int

func (v value) neg() value {
	res := make(value)
	for key, coeff := range v {
		res[key] = -coeff
	}
	return res
}

func (v value) mul(b value) value {
	if v == nil {
		return b
	}
	if b == nil {
		return v
	}
	res := make(value)
	for v, coef := range v {
		for v2, coef2 := range b {
			c := coef * coef2
			f := v.mul(v2)
			res[f] += c
		}
	}
	return res
}

func (v value) add(b value) value {
	if v == nil {
		return b
	}
	if b == nil {
		return v
	}

	res := make(value)
	for s, x := range v {
		res[sortstr(s)] += x
	}
	for s, x := range b {
		res[sortstr(s)] += x
	}
	return res
}

const mul = 0
const add = 1
const sub = 2

func (p *parser) parseExpr() value {
	v := p.parseTerm()
	for p.i < len(p.fields) {
		switch p.fields[p.i] {
		case "+":
			p.i++
			v2 := p.parseTerm()
			v = v.add(v2)
		case "-":
			p.i++
			v2 := p.parseTerm()
			v = v.add(v2.neg())
		default:
			return v
		}
	}
	return v
}

func (p *parser) parseTerm() value {
	v := p.parseFactor()
	for p.i < len(p.fields) {
		switch p.fields[p.i] {
		case "*":
			p.i++
			v2 := p.parseFactor()
			v = v.mul(v2)
		default:
			return v
		}
	}
	return v
}

func (p *parser) parseFactor() value {
	if p.fields[p.i] == "(" {
		p.i++
		v := p.parseExpr()
		p.i++ // ")"
		return v
	}
	if unicode.IsLetter(rune(p.fields[p.i][0])) {
		f := p.fields[p.i]
		if v, exists := p.m[f]; exists {
			p.i++
			return value{"": v}
		}
		v := value{facstr(p.fields[p.i]): 1}
		p.i++
		return v
	}
	var x int
	for i := range p.fields[p.i] {
		x *= 10
		x += int(p.fields[p.i][i] - '0')
	}
	p.i++
	return value{"": x}
}

type parser struct {
	fields []string
	m      map[string]int
	i      int
}

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	// Replace vars with values
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")
	fields := strings.Fields(expression)

	m := make(map[string]int)
	for i := range evalvars {
		m[evalvars[i]] = evalints[i]
	}
	var p parser
	p.fields = fields
	p.m = m
	val := p.parseExpr()

	haha := make(value)
	for k, v := range val {
		k = sortstr(k)
		haha[k] += v
	}
	val = haha

	type kv struct {
		coeff int
		vars  []string
	}
	var kvs []kv
	for k, v := range val {
		kvs = append(kvs, kv{v, strings.Split(string(k), "*")})
		sort.Strings(kvs[len(kvs)-1].vars)
	}
	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].vars[0] == "" {
			return false
		}
		if kvs[j].vars[0] == "" {
			return true
		}
		if len(kvs[i].vars) == len(kvs[j].vars) {
			for k := range kvs[i].vars {
				if kvs[i].vars[k] < kvs[j].vars[k] {
					return true
				}
				if kvs[i].vars[k] > kvs[j].vars[k] {
					return false
				}
			}
			return true
		}
		return len(kvs[i].vars) > len(kvs[j].vars)
	})

	var res []string
	for _, s := range kvs {
		if s.coeff == 0 {
			continue
		}
		if len(s.vars) == 0 || s.vars[0] == "" {
			res = append(res, fmt.Sprint(s.coeff))
		} else {
			res = append(res, fmt.Sprintf("%v*%v", s.coeff, strings.Join(s.vars, "*")))
		}
	}

	return res
}

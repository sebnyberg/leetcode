package p0726numberofatoms

import (
	"fmt"
	"sort"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_countOfAtoms(t *testing.T) {
	for _, tc := range []struct {
		formula string
		want    string
	}{
		{"H2O", "H2O"},
		{"Mg(OH)2", "H2MgO2"},
		{"(OH)12", "H12O12"},
		{"K4(ON(SO3)2)2", "K4N2O14S4"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.formula), func(t *testing.T) {
			require.Equal(t, tc.want, countOfAtoms(tc.formula))
		})
	}
}

func countOfAtoms(formula string) string {
	// When encountering '(', start a new stack
	// When encountering a letter, add next digit and add to current level in
	// stack.
	// When encountering ')' read next digit and multiply current stack before
	// merging with previous level.
	atomStack := []map[string]int{{}}
	parseDigit := func(s string) (x int, n int, ok bool) {
		if len(s) == 0 || s[0] < '0' || s[0] > '9' {
			return 0, 0, false
		}
		var i int
		for i = 0; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			x *= 10
			x += int(s[i] - '0')
		}
		return x, i, true
	}

	var pos int
	for pos < len(formula) {
		switch {
		case formula[pos] == '(':
			pos++
			atomStack = append(atomStack, map[string]int{})
		case formula[pos] == ')':
			pos++
			a := atomStack[len(atomStack)-1]
			x, n, ok := parseDigit(formula[pos:])
			if ok {
				pos += n
				// Multiply contents of stack with digit
				for k := range a {
					a[k] *= x
				}
			}
			// Merge with previous level
			atomStack = atomStack[:len(atomStack)-1] // pop
			for k := range a {
				atomStack[len(atomStack)-1][k] += a[k]
			}
		case formula[pos] >= 'A' && formula[pos] <= 'Z':
			elem := []byte{formula[pos]}
			pos++
			for pos < len(formula) && unicode.IsLower(rune(formula[pos])) {
				elem = append(elem, formula[pos])
				pos++
			}
			x, n, ok := parseDigit(formula[pos:])
			if ok {
				pos += n
				atomStack[len(atomStack)-1][string(elem)] += x
			} else {
				atomStack[len(atomStack)-1][string(elem)]++
			}
		}
	}
	res := make([]byte, 0, 26)
	elems := make([]string, len(atomStack[0]))
	for k := range atomStack[0] {
		elems = append(elems, k)
	}
	sort.Strings(elems)
	for _, elem := range elems {
		count := atomStack[0][elem]
		if count == 0 {
			continue
		}
		res = append(res, elem...)
		if count > 1 {
			res = append(res, fmt.Sprint(count)...)
		}
	}
	return string(res)
}

package p0816ambiguouscoordinates

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ambiguousCoordinates(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{"(123)", []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"}},
		{"(00011)", []string{"(0.001, 1)", "(0, 0.011)"}},
		{"(0123)", []string{"(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"}},
		{"(100)", []string{"(10, 0)"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, ambiguousCoordinates(tc.s))
		})
	}
}

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	// There are two basic operations:
	// 1. Splitting the string into two components
	// 2. Placing a single decimal somewhere in a component
	aComps := make([]string, 0)
	bComps := make([]string, 0)
	res := make([]string, 0)
	for i := 1; i < len(s); i++ {
		// Split on i, forming two components
		a, b := s[:i], s[i:]
		aComps = aComps[:0]
		bComps = bComps[:0]
		if isValid(a) {
			aComps = append(aComps, a)
		}
		if isValid(b) {
			bComps = append(bComps, b)
		}

		// For each 'a' component, insert a dot in each possible position,
		// adding the result to buf
		for i := 1; i < len(a); i++ {
			withDot := a[:i] + "." + a[i:]
			if !isValid(withDot) {
				continue
			}
			aComps = append(aComps, withDot)
		}
		if len(aComps) == 0 {
			continue
		}

		// For each 'b' component, insert a dot in each possible position,
		// adding the result to buf
		for i := 1; i < len(b); i++ {
			withDot := b[:i] + "." + b[i:]
			if !isValid(withDot) {
				continue
			}
			bComps = append(bComps, withDot)
		}
		if len(bComps) == 0 {
			continue
		}

		// Finally, for each combination, add to the result and continue
		for _, a := range aComps {
			for _, b := range bComps {
				res = append(res, "("+a+", "+b+")")
			}
		}
	}
	return res
}

func isValid(p string) bool {
	if strings.ContainsAny(p, ".") && strings.HasSuffix(p, "0") {
		return false
	}
	if len(p) > 1 && p[1] != '.' && p[0] == '0' {
		return false
	}
	return true
}

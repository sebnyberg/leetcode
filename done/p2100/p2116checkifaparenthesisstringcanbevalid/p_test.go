package p2116checkifaparenthesisstringcanbevalid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canBeValid(t *testing.T) {
	for _, tc := range []struct {
		s      string
		locked string
		want   bool
	}{
		{"()))(()(()()()()(((())())((()((())", "1100000000000010000100001000001101", true},
		{"()", "11", true},
		{"))()))", "010100", true},
		{"()()", "0000", true},
		{")", "0", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, canBeValid(tc.s, tc.locked))
		})
	}
}

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 { // must be even
		return false
	}

	// First, iterate through s and check if any of the locked parenthesis can
	// be skipped
	stack := []int{}
	skip := make([]bool, len(s))
	for i := range s {
		if locked[i] == '0' {
			continue
		}
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				// Pop and mark the pair as 'skip'
				l := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				skip[l] = true
				skip[i] = true
			}
		}
	}

	// The first parenthesis must be open.
	// The second can be either open or closed, it depends on other parenthesis

	// When encountering a locked parenthesis pointing to the left, there must
	lock := make([]bool, len(s))
	available := make([]bool, len(s))
	for i, ch := range locked {
		if skip[i] {
			continue
		}
		if ch == '1' {
			lock[i] = true
		} else {
			available[i] = true
		}
	}
	// Scan from left to right, checking left locked parenthesis
	l := 0
	for i := range s {
		if skip[i] || !lock[i] || s[i] != ')' {
			continue
		}

		// If the position is locked and facing left, check if there is an available
		// parenthesis on the left that can balance out the expression
		for ; l < i && (!available[l] || skip[l]); l++ {
		}
		if l == i { // no available parenthesis to match this one with
			return false
		}
		available[l] = false
	}

	// Scan from right to left, checking right locked parenthesis
	r := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if skip[i] || !lock[i] || s[i] != '(' {
			continue
		}
		// If the position is locked and facing left, check if there is an available
		// parenthesis on the left that can balance out the expression
		for ; r != i && (!available[r] || skip[r]); r-- {
		}
		if r == i { // no available parenthesis to match this one with
			return false
		}
		available[r] = false
	}
	return true
}

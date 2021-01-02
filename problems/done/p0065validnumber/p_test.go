package p0064validnumber

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_isNumber(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want bool
	}{
		// {" ", false},
		// {"0", true},
		{" 0.1", true},
		// {" 3.", true},
		// {"abc", false},
		// {"1 a", false},
		// {"2e10", true},
		// {" -90e3", true},
		// {" 1e", false},
		// {"e3", false},
		// {" 6e-1", true},
		// {" 99e2.5", false},
		// {"53.5e93", true},
		// {" --6", false},
		// {"-+3", false},
		// {"95a54e53", false},
		// {" .1", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, isNumber(tc.in))
		})
	}
}

type validator struct {
	in    string
	pos   int
	width int
}

var eof = rune(0)

func (v *validator) next() (ch rune) {
	if v.pos >= len(v.in) {
		v.width = 0
		return eof
	}
	ch, v.width = utf8.DecodeRuneInString(v.in[v.pos:])
	v.pos += v.width
	return ch
}

func (v *validator) accept(valid string) bool {
	if strings.ContainsRune(valid, v.next()) {
		return true
	}
	v.backup()
	return false
}

func (v *validator) acceptRun(valid string) (found bool) {
	for strings.ContainsRune(valid, v.next()) {
		found = true
	}
	v.backup()
	return found
}

func (v *validator) backup() {
	v.pos -= v.width
}

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if strings.ContainsRune(s, ' ') {
		return false
	}
	v := validator{in: s}
	digits := "0123456789"
	// Optional leading sign
	v.accept("+-")
	pre := v.acceptRun(digits)
	if v.accept(".") {
		if !v.acceptRun(digits) && !pre {
			return false
		}
	} else {
		if !pre {
			return false
		}
	}
	if v.accept("eE") {
		v.accept("+-")
		if !v.acceptRun(digits) {
			return false
		}
	}
	return v.next() == eof
}

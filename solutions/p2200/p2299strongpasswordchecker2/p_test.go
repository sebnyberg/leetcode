package p2299strongpasswordchecker2

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_strongPasswordCheckerII(t *testing.T) {
	for _, tc := range []struct {
		password string
		want     bool
	}{
		{"11A!A!Aa", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.password), func(t *testing.T) {
			require.Equal(t, tc.want, strongPasswordCheckerII(tc.password))
		})
	}
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	var hasLower, hasUpper, hasDigit bool
	if !strings.ContainsAny(password, "!@#$%^&*()-+") {
		return false
	}
	for i, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = true
		}
		if unicode.IsUpper(ch) {
			hasUpper = true
		}
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}
	return hasLower && hasUpper && hasDigit
}

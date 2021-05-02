package p1849splittingastringintodescendingconsecutivevalues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"20000000000000000001", true},
		{"050043", true},
		{"1234", false},
		{"9080701", false},
		{"10009998", true},
		{"14188802907687215148", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, splitString(tc.s))
		})
	}
}

func splitString(s string) bool {
	return helper(s, 0, 0)
}

func helper(s string, count int, prev int) bool {
	if len(s) == 0 && count >= 2 {
		return true
	}

	for i := 1; i <= len(s); i++ {
		n := parseStr(s[:i])
		if count > 0 && n != prev-1 {
			continue
		}
		if helper(s[i:], count+1, n) {
			return true
		}
	}
	return false
}

func parseStr(s string) int {
	var res int
	for i := range s {
		res *= 10
		res += int(s[i] - '0')
	}
	return res
}

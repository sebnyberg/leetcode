package p0263uglynumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isUgly(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{6, true},
		{8, true},
		{14, false},
		{1, true},
		{-2147483648, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isUgly(tc.n))
		})
	}
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

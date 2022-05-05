package p0202happynumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isHappy(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isHappy(tc.n))
		})
	}
}

func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for {
		next := 0
		for n > 0 {
			next += (n % 10) * (n % 10)
			n /= 10
		}
		if _, ok := seen[next]; ok {
			return false
		}
		if next == 1 {
			return true
		}
		seen[next] = struct{}{}
		n = next
	}
}

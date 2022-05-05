package p1952threedivisors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isThree(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{2, false},
		{4, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isThree(tc.n))
		})
	}
}

func isThree(n int) bool {
	// All integers are divisible by 1 and itself, so we want to find if the
	// number is prime. Since it's only 10^4, we can brute-force without sieve.
	var found bool
	for a := 2; a < n; a++ {
		if n%a == 0 {
			if found {
				return false
			}
			found = true
		}
	}
	return found
}

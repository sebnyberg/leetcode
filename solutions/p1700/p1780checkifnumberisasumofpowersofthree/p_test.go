package p1780checkifnumberisasumofpowersofthree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPowersOfThree(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{12, true},
		{91, true},
		{21, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, checkPowersOfThree(tc.n))
		})
	}
}

func checkPowersOfThree(n int) bool {
	// All powers of three are divisible by three except when
	// combined with 3^0, in which case n%3 == 1
	if n%3 == 2 {
		return false
	}
	if n%3 == 1 {
		n--
	}
	// We know now that n is divisible by 3, but not necessarily if it
	// is a sum of distinct powers of three
	// Let's try a recursion brute-force:
	return findPowersOfThree(n, 1, 3)
}

func findPowersOfThree(n, power, factor int) bool {
	if n < 0 || power >= 16 {
		return false
	} else if n == 0 {
		return true
	}

	return findPowersOfThree(n, power+1, factor*3) ||
		findPowersOfThree(n-factor, power+1, factor*3)
}

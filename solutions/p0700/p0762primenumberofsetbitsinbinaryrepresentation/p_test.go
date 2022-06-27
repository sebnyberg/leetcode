package p0762primenumberofsetbitsinbinaryrepresentation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPrimeSetBits(t *testing.T) {
	for _, tc := range []struct {
		left  int
		right int
		want  int
	}{
		{6, 10, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, countPrimeSetBits(tc.left, tc.right))
		})
	}
}

func countPrimeSetBits(left int, right int) int {
	notPrime := make([]bool, 33)
	notPrime[0] = true
	notPrime[1] = true
	for i := 2; i <= 32; i++ {
		for j := 2 * i; j <= 32; j += i {
			notPrime[j] = true
		}
	}

	hasPrimeBits := func(x int) bool {
		var count int
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return !notPrime[count]
	}
	var res int
	for x := left; x <= right; x++ {
		if hasPrimeBits(x) {
			res++
		}
	}
	return res
}

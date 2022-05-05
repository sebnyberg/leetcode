package p0343integerbreak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integerBreak(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 1},
		{10, 36},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, integerBreak(tc.n))
		})
	}
}

func integerBreak(n int) int {
	// There are four cases:
	// n <= 3 => answer is n-1
	// n % 3 == 0 => answer is 3^(n/3)
	// n % 3 == 1 => answer is 3^((n/3)-1)*4
	// n % 3 == 2 => answer is 3^(n/3)*2
	if n <= 3 {
		return n - 1
	}
	p := n / 3
	switch n % 3 {
	case 0:
		return pow(3, p)
	case 1:
		return pow(3, p-1) * 4
	case 2:
		return pow(3, p) * 2
	}
	return -1
}

func pow(x, p int) int {
	if p == 0 {
		return 1
	}
	res := pow(x, p/2)
	if p%2 == 1 {
		return res * res * x
	} else {
		return res * res
	}
}

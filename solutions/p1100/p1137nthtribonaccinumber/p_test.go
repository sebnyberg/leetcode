package p1137nthtribonaccinumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tribonacci(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{4, 4},
		{25, 1389537},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, tribonacci(tc.n))
		})
	}
}

func tribonacci(n int) int {
	fib := [3]int{0, 1, 1}
	for n >= 3 {
		fib[0], fib[1], fib[2] = fib[1], fib[2], fib[0]+fib[1]+fib[2]
		n--
	}
	return fib[n]
}

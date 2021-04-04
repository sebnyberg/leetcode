package p0650twokeyskeyboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSteps(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1000, 21},
		{3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minSteps(tc.n))
		})
	}
}

// Todo: prime factorization - why does it work?
func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	return 1 + dp(map[[3]int]int{}, n, 1, 1)
}

func dp(mem map[[3]int]int, n, m, buflen int) int {
	if m == n {
		return 0
	}
	if m > n {
		return 1000
	}
	k := [3]int{n, m, buflen}
	if v, exists := mem[k]; exists {
		return v
	}
	res := min(
		2+dp(mem, n, m*2, m),           // copy + paste
		1+dp(mem, n, m+buflen, buflen), // paste
	)
	mem[k] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

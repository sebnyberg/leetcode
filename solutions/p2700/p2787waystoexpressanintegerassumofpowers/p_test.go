package p2787waystoexpressanintegerassumofpowers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWays(t *testing.T) {
	for i, tc := range []struct {
		n    int
		x    int
		want int
	}{
		{10, 2, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWays(tc.n, tc.x))
		})
	}
}

func numberOfWays(n int, x int) int {
	mem := make(map[[2]int]int)
	res := f(mem, n, x, 1)
	return res % mod
}

const mod = 1e9 + 7

func f(mem map[[2]int]int, n, x, k int) int {
	key := [2]int{n, k}
	if v, exists := mem[key]; exists {
		return v
	}
	if n == 0 {
		return 1
	}
	if pow(k, x) > n {
		return 0
	}
	res := f(mem, n-pow(k, x), x, k+1) + f(mem, n, x, k+1)
	mem[key] = res % mod
	return res
}

func pow(a, b int) int {
	res := a
	for i := 0; i < b-1; i++ {
		res *= a
	}
	return res
}

package p0264uglynumber2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthUglyNumber(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{10, 12},
		{11, 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nthUglyNumber(tc.n))
		})
	}
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	k := make([]int, 1)
	k[0] = 1
	t1, t2, t3 := 0, 0, 0
	for i := 1; i < n; i++ {
		k1, k2, k3 := k[t1]*2, k[t2]*3, k[t3]*5
		k = append(k, min(k1, min(k2, k3)))
		if k[i] == k1 {
			t1++
		}
		if k[i] == k2 {
			t2++
		}
		if k[i] == k3 {
			t3++
		}
	}
	return k[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

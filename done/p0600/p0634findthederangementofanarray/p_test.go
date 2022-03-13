package p0634findthederangementofanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDerangement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{3, 2},
		{2, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findDerangement(tc.n))
		})
	}
}

const mod = 1e9 + 7

func findDerangement(n int) int {
	if n == 1 {
		return 0
	}
	res := 1
	for x := 1; x <= n; x++ {
		if x%2 == 0 {
			res = (res*x + 1) % mod
		} else {
			res = (res*x - 1) % mod
		}
	}
	return res
}

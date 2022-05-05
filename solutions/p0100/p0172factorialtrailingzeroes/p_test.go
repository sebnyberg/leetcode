package p0172factorialtrailingzeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{30, 7},
		{3, 0},
		{5, 1},
		{0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, trailingZeroes(tc.n))
		})
	}
}
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n /= 5
	}
	return res
}

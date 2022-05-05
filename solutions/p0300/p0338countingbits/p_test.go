package p0338countingbits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countBits(tc.n))
		})
	}
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

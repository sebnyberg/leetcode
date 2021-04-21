package p0276paintfence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numWays(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{3, 2, 6},
		{1, 1, 1},
		{7, 2, 42},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, numWays(tc.n, tc.k))
		})
	}
}

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	same := 0
	different := k
	for i := 2; i <= n; i++ {
		tmp := same
		same = different
		different = (tmp + different) * (k - 1)
	}
	return same + different
}

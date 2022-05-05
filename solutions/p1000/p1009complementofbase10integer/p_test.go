package p1009complementofbase10integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bitwiseComplement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{5, 2},
		{7, 0},
		{10, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, bitwiseComplement(tc.n))
		})
	}
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	var mask int
	for m := n; m > 0; m >>= 1 {
		mask <<= 1
		mask |= 1
	}
	res := ^n & mask
	return res
}

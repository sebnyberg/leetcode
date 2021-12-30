package p1015smallestintegerdivisiblebyk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestRepunitDivByK(t *testing.T) {
	for _, tc := range []struct {
		k    int
		want int
	}{
		{1, 1},
		{2, -1},
		{3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, smallestRepunitDivByK(tc.k))
		})
	}
}

func smallestRepunitDivByK(k int) int {
	// No odd number is evenly divisible with 2 and 5
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	r := 0
	for n := 1; n <= k; n++ {
		r = (r*10 + 1) % k
		if r == 0 {
			return n
		}
	}
	return -1
}

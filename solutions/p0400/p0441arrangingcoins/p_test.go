package p0441arrangingcoins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_arrangeCoins(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{5, 2},
		{8, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, arrangeCoins(tc.n))
		})
	}
}

func arrangeCoins(n int) int {
	var val, i int
	for n >= val {
		i++
		val += i
	}
	return i - 1
}

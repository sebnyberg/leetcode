package p0476numbercomplement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findComplement(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{5, 2},
		{1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, findComplement(tc.num))
		})
	}
}

func findComplement(num int) int {
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	mask--
	return ^num & mask
}

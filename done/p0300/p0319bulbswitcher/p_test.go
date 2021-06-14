package p0319bulbswitcher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bulbSwitch(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{3, 1},
		{0, 0},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, bulbSwitch(tc.n))
		})
	}
}

func bulbSwitch(n int) int {
	// All divisors come in pairs (1, 12), (2,6), (3,4) except numbers which are
	// squares, they "flip the switch" an extra time. Thus, the solution is to
	// count the nubmer of squares from 1 to n
	var count int
	for i := 1; i*i <= n; i++ {
		count++
	}
	return count
}

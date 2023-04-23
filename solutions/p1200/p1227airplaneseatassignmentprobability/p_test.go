package p1227airplaneseatassignmentprobability

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthPersonGetsNthSeat(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want float64
	}{
		{2, .5},
		{1, 1.0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := nthPersonGetsNthSeat(tc.n)
			require.InEpsilon(t, tc.want, res, 0.001)
		})
	}
}

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}

package p2119anumberafterdoublereversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSameAfterReversals(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{526, true},
		{1800, false},
		{0, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, isSameAfterReversals(tc.num))
		})
	}
}

func isSameAfterReversals(num int) bool {
	// The requirement is true as long as there are no leading zeroes
	return num == 0 || num%10 != 0
}

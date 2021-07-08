package p0367validperfectsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPerfectSquare(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{16, true},
		{14, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, isPerfectSquare(tc.num))
		})
	}
}

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

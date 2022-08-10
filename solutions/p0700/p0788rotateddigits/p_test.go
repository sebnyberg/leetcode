package p0788rotateddigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotatedDigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{857, 247},
		{10, 4},
		{1, 0},
		{2, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, rotatedDigits(tc.n))
		})
	}
}

func rotatedDigits(n int) int {
	var count int
	for x := 1; x <= n; x++ {
		var good bool
		var bad bool
		for y := x; y > 0; y /= 10 {
			m := y % 10
			if m == 2 || m == 5 || m == 6 || m == 9 {
				good = true
			}
			if m == 3 || m == 4 || m == 7 {
				bad = true
			}
		}
		if !bad && good {
			count++
		}
	}
	return count
}

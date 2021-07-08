package p0357countnumberswithuniquedigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 91},
		{0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countNumbersWithUniqueDigits(tc.n))
		})
	}
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	var counts int
	if n >= 1 {
		counts = 10
	}
	for i := 2; i <= n; i++ {
		res := 9
		cur := 9
		for j := i - 1; j > 0; j-- {
			res *= cur
			cur--
		}
		counts += res
	}
	return counts
}

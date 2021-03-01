package p0233numdigitone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countDigitOne(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{20, 12},
		{13, 6},
		{1, 1},
		{11, 4},
		{100, 21},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countDigitOne(tc.n))
		})
	}
}

func countDigitOne(n int) (res int) {
	i := 0
	divisor := 1
	for n >= divisor {
		nextdivisor := divisor * 10
		res += (n / nextdivisor) * divisor
		rest := n % nextdivisor
		if rest >= divisor && rest < 2*divisor {
			res += (n % divisor) + 1
		} else if rest >= 2*divisor {
			res += divisor
		}
		i++
		divisor = nextdivisor
	}

	return res
}

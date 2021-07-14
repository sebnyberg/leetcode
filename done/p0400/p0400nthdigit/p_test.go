package p0400nthdigit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findNthDigit(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{9, 9},
		{11, 0},
		{3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findNthDigit(tc.n))
		})
	}
}

func findNthDigit(n int) int {
	n--
	ndigit := 1
	floor := 1
	ceil := 10
	for n >= ndigit*(ceil-floor) {
		n -= (ceil - floor) * ndigit
		ceil *= 10
		floor *= 10
		ndigit++
	}
	num := floor + n/ndigit
	idxInNum := n % ndigit
	for i := 0; i < ndigit-1-idxInNum; i++ {
		num /= 10
	}
	res := num % 10
	return res
}

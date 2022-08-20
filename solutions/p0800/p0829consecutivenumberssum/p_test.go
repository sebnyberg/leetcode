package p0829consecutivenumberssum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_consecutiveNumbersSum(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{85, 4},
		{15, 4},
		{9, 3},
		{5, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, consecutiveNumbersSum(tc.n))
		})
	}
}

func consecutiveNumbersSum(n int) int {
	res := 1
	minSum := 1
	for x := 2; ; x++ {
		minSum += x
		if n < minSum {
			break
		}
		if x&1 == 1 {
			if n%x == 0 {
				res++
			}
		} else {
			if (n-(x/2))%x == 0 {
				res++
			}
		}
	}
	return res
}

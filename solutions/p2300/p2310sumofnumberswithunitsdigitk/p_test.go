package p2310sumofnumberswithunitsdigitk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums int
		k    int
		want int
	}{
		{10, 1, 10},
		{2, 8, -1},
		{58, 9, 2},
		{37, 2, -1},
		{0, 7, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumNumbers(tc.nums, tc.k))
		})
	}
}

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for factor := 1; factor <= 10; factor++ {
		if (k*factor)%10 == num%10 && k*factor <= num {
			return factor
		}
	}
	return -1
}

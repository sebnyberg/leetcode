package p2169countoperationstoobtainzero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countOperations(t *testing.T) {
	for _, tc := range []struct {
		num1 int
		num2 int
		want int
	}{
		{2, 3, 3},
		{10, 10, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num1), func(t *testing.T) {
			require.Equal(t, tc.want, countOperations(tc.num1, tc.num2))
		})
	}
}

func countOperations(num1 int, num2 int) int {
	var ops int
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ops++
	}
	return ops
}

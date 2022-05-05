package p0991brokencalculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_brokenCalc(t *testing.T) {
	for _, tc := range []struct {
		startValue int
		target     int
		want       int
	}{
		{1, 1e9, 39},
		{1024, 1, 1023},
		{2, 3, 2},
		{5, 8, 2},
		{3, 10, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startValue), func(t *testing.T) {
			require.Equal(t, tc.want, brokenCalc(tc.startValue, tc.target))
		})
	}
}

func brokenCalc(startValue int, target int) int {
	var steps int
	for target > startValue {
		steps++
		if target&1 == 1 {
			target++
		} else {
			target >>= 1
		}
	}
	return steps + (startValue - target)
}

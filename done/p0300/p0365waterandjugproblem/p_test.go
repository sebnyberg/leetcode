package p0365waterandjugproblem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canMeasureWater(t *testing.T) {
	for _, tc := range []struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
		want           bool
	}{
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.jug1Capacity), func(t *testing.T) {
			require.Equal(t, tc.want, canMeasureWater(tc.jug1Capacity, tc.jug2Capacity, tc.targetCapacity))
		})
	}
}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}
	for jug2Capacity != 0 {
		jug2Capacity, jug1Capacity = jug1Capacity%jug2Capacity, jug2Capacity
	}
	return targetCapacity%jug1Capacity == 0
}

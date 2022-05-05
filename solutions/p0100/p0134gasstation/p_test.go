package p0134gasstation

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canCompleteCircuit(t *testing.T) {
	for _, tc := range []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{3, 1, 1}, []int{1, 2, 2}, 0},
		{[]int{4, 5, 3, 1, 4}, []int{5, 4, 3, 4, 2}, -1},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.gas, tc.cost), func(t *testing.T) {
			require.Equal(t, tc.want, canCompleteCircuit(tc.gas, tc.cost))
		})
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	// Deduct the cost from each gas station
	// If the sum of the gas difference is smaller than zero, there is no
	// way to complete the lap.
	// If there is a solution (sum > 0, and there is only one per problem desc.)
	// it must be after the most negative net gas value at any station.
	var sum int
	minVal := math.MaxInt32
	var minIndex int
	for i := range gas {
		gas[i] -= cost[i]
		sum += gas[i]
		if sum < minVal {
			minVal = sum
			minIndex = i
		}
	}
	if sum < 0 {
		return -1
	}
	return (minIndex + 1) % len(gas)
}

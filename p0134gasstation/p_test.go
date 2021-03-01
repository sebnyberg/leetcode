package p0134gasstation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canCompleteCircuit(t *testing.T) {
	for _, tc := range []struct {
		gas  []int
		cost []int
		want int
	}{
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
	for i := range gas {
		var tank int
		var j int
		for j = 0; ; j++ {
			pos := (i + j) % len(gas)
			tank += gas[pos]
			tank -= cost[pos]
			if tank <= 0 || j == len(gas)-1 {
				break
			}
		}
		if tank > 0 || (j == len(gas)-1 && tank == 0) {
			return i
		}
	}
	return -1
}

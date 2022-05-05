package p2105wateringplants2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumRefill(t *testing.T) {
	for _, tc := range []struct {
		plants               []int
		capacityA, capacityB int
		want                 int
	}{
		{[]int{2, 2, 3, 3}, 5, 5, 1},
		{[]int{2, 2, 3, 3}, 3, 4, 2},
		{[]int{5}, 10, 8, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.plants), func(t *testing.T) {
			require.Equal(t, tc.want, minimumRefill(tc.plants, tc.capacityA, tc.capacityB))
		})
	}
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	alice, bob := 0, len(plants)-1
	aliceWater, bobWater := capacityA, capacityB
	var refill int
	for alice < bob {
		if aliceWater < plants[alice] {
			aliceWater = capacityA
			refill++
		}
		aliceWater -= plants[alice]
		if bobWater < plants[bob] {
			bobWater = capacityB
			refill++
		}
		bobWater -= plants[bob]
		alice++
		bob--
	}
	if alice == bob && max(aliceWater, bobWater) < plants[alice] {
		refill++
	}
	return refill
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2079wateringplants

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wateringPlants(t *testing.T) {
	for _, tc := range []struct {
		plants   []int
		capacity int
		want     int
	}{
		{[]int{2, 2, 3, 3}, 5, 14},
		{[]int{1, 1, 1, 4, 2, 3}, 4, 30},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 8, 49},
	} {
		t.Run(fmt.Sprintf("%+v", tc.plants), func(t *testing.T) {
			require.Equal(t, tc.want, wateringPlants(tc.plants, tc.capacity))
		})
	}
}

func wateringPlants(plants []int, capacity int) int {
	// Iterating is still O(n) so should be fine for 1e6
	steps := 1
	water := capacity - plants[0]
	for i := 0; i < len(plants)-1; i++ {
		if plants[i+1] > water {
			steps += i + 1
			steps += i + 1
			water = capacity
		}
		steps++
		water -= plants[i+1]
	}
	return steps
}

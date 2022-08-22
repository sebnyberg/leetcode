package p0853carfleet

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_carFleet(t *testing.T) {
	for _, tc := range []struct {
		target   int
		position []int
		speed    []int
		want     int
	}{
		{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, carFleet(tc.target, tc.position, tc.speed))
		})
	}
}

func carFleet(target int, position []int, speed []int) int {
	// Let's consider cars from the last to first car
	// If the current car is faster than the previous car, then we increment the
	// result by 1.
	// If it is not faster than the previous car, then we must check whether the
	// previous car will bump into the current car, and at what time that will
	// happen. If it happens before the cars have reached the end, then we
	// continue to the previous car.
	//
	type item struct {
		pos   int
		speed int
	}
	n := len(position)
	items := make([]item, n)
	for i := range items {
		items[i] = item{position[i], speed[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].pos < items[j].pos
	})
	res := 1
	j := n - 1
	for i := n - 2; i >= 0; i-- {
		// Check if this car will bump into the next car
		if items[i].speed < items[j].speed {
			j = i
			res++
			continue
		}
		dx1 := target - items[j].pos
		t1 := float64(dx1) / float64(items[j].speed)
		dx2 := target - items[i].pos
		t2 := float64(dx2) / float64(items[i].speed)
		if t1 < t2 {
			j = i
			res++
			continue
		}
	}
	return res
}

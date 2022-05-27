package p1552magneticofrcebetweentwoballs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDistance(t *testing.T) {
	for _, tc := range []struct {
		position []int
		m        int
		want     int
	}{
		{[]int{1, 2, 3, 4, 7}, 3, 3},
		{[]int{5, 4, 3, 2, 1, 1e9}, 2, 1e9 - 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.position), func(t *testing.T) {
			require.Equal(t, tc.want, maxDistance(tc.position, tc.m))
		})
	}
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	var max int
	for _, n := range position {
		if n > max {
			max = n
		}
	}

	// canAchieveDist returns true if it is possible to achieve a minimum distance
	// of dist when distributing the balls in the baskets.
	canAchieveDist := func(dist int) bool {
		mm := m
		mm--
		prev := position[0]
		for i := 1; mm > 0 && i < len(position); i++ {
			d := position[i] - prev
			if d >= dist {
				prev = position[i]
				mm--
			}
		}
		return mm == 0
	}

	lo, hi := 0, max+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if canAchieveDist(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

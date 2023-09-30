package p2751robotcollisions

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_survivedRobotsHealths(t *testing.T) {
	for i, tc := range []struct {
		positions  []int
		healths    []int
		directions string
		want       []int
	}{
		{[]int{1, 40}, []int{10, 11}, "RL", []int{10}},
		{[]int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL", []int{14}},
		{[]int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR", []int{2, 17, 9, 15, 10}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, survivedRobotsHealths(tc.positions, tc.healths, tc.directions))
		})
	}
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	// Only robots that are moving towards each other can collide.
	// In fact, I think we can create a stack of robots moving to the right and
	// pop whenever a robot collides with another robot on the stack. Also, I
	// don't even think time matters, only direction.
	//

	// First, sort by position
	n := len(positions)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return positions[idx[i]] < positions[idx[j]]
	})
	res := make([]int, n)

	healthRight := []int{}
	for _, j := range idx {
		if directions[j] == 'R' {
			healthRight = append(healthRight, j)
			continue
		}
		for len(healthRight) > 0 && healths[healthRight[len(healthRight)-1]] <= healths[j] {
			healths[j]--
			healthRight = healthRight[:len(healthRight)-1]
		}
		if len(healthRight) > 0 && healths[healthRight[len(healthRight)-1]] == healths[j] {
			healthRight = healthRight[:len(healthRight)-1]
			continue
		}
		if len(healthRight) > 0 {
			healths[healthRight[len(healthRight)-1]]--
			continue
		}
		if healths[j] > 0 {
			// robot will keep going left forever
			res[j] = healths[j]
		}
	}
	for _, h := range healthRight {
		res[h] = healths[h]
	}
	var j int
	for i := range res {
		if res[i] != 0 {
			res[j] = res[i]
			j++
		}
	}
	res = res[:j]
	return res
}

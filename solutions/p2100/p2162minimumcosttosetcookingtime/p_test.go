package p2162minimumcosttosetcookingtime

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostSetTime(t *testing.T) {
	for _, tc := range []struct {
		startAt       int
		moveCost      int
		pushCost      int
		targetSeconds int
		want          int
	}{
		{0, 1, 1, 0, 0},
		{1, 2, 1, 600, 6},
		{0, 1, 2, 76, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startAt), func(t *testing.T) {
			require.Equal(t, tc.want, minCostSetTime(tc.startAt, tc.moveCost, tc.pushCost, tc.targetSeconds))
		})
	}
}

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	// At least 1 second
	// At most 99 minutes and 99 seconds
	type alternative struct {
		minutes int
		seconds int
	}
	alternatives := make([]alternative, 0)
	needMinutes := targetSeconds / 60
	secondsLeft := targetSeconds % 60
	if needMinutes < 100 {
		// Easy scenario
		alternatives = append(alternatives, alternative{
			minutes: needMinutes,
			seconds: secondsLeft,
		})
	}
	// Move 1 minute to seconds
	if needMinutes >= 1 && secondsLeft <= 39 {
		alternatives = append(alternatives, alternative{
			minutes: needMinutes - 1,
			seconds: secondsLeft + 60,
		})
	}

	str := func(a alternative) string {
		if a.minutes > 0 {
			return fmt.Sprintf("%d%02d", a.minutes, a.seconds)
		}
		return fmt.Sprintf("%d", a.seconds)
	}

	// The moves are symmetric, i.e. it doesn't matter if we start in position
	// 1 and move left, then right, or at second to last, move right, then left
	// Also it does not matter if we push in the first place - if we have to move
	// then we might as well do it immediately
	res := math.MaxInt32
	for _, alt := range alternatives {
		s := str(alt)
		cur := startAt
		var cost int
		for _, ch := range s {
			x := int(ch - '0')
			if x != cur {
				cur = x
				cost += moveCost
			}
			cost += pushCost
		}
		res = min(res, cost)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

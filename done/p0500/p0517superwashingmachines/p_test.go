package p0517superwashingmachines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinMoves(t *testing.T) {
	for _, tc := range []struct {
		machines []int
		want     int
	}{
		{[]int{1, 0, 5}, 3},
		{[]int{0, 3, 0}, 2},
		{[]int{0, 2, 0}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.machines), func(t *testing.T) {
			require.Equal(t, tc.want, findMinMoves(tc.machines))
		})
	}
}

func findMinMoves(machines []int) int {
	n := len(machines)
	pre := make([]int, n+1)
	for i := range machines {
		pre[i+1] = pre[i] + machines[i]
	}
	total := pre[n]
	if total%n != 0 {
		return -1
	}
	sz := total / n

	var res int
	for i := range machines {
		left := pre[i]
		right := pre[n] - pre[i+1]
		wantLeft := i * sz
		wantRight := (n - i - 1) * sz
		var leftMoves int
		if left < wantLeft {
			leftMoves = wantLeft - left
		}
		var rightMoves int
		if right < wantRight {
			rightMoves = wantRight - right
		}
		res = max(res, leftMoves+rightMoves)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

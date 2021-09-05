package p1936addminimumnumberofrungs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addRungs(t *testing.T) {
	for _, tc := range []struct {
		rungs []int
		dist  int
		want  int
	}{
		{[]int{3}, 1, 2},
		{[]int{1, 3, 5, 10}, 2, 2},
		{[]int{3, 6, 8, 10}, 3, 0},
		{[]int{3, 4, 6, 7}, 2, 1},
		{[]int{5}, 10, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rungs), func(t *testing.T) {
			require.Equal(t, tc.want, addRungs(tc.rungs, tc.dist))
		})
	}
}

func addRungs(rungs []int, dist int) int {
	var cur int
	var i int
	var extraRungs int
	for i < len(rungs) {
		delta := rungs[i] - cur
		if delta <= dist {
			cur = rungs[i]
			i++
			continue
		}
		jumps := max(1, (delta/dist)-1)
		extraRungs += jumps
		cur += dist * jumps
	}
	return extraRungs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

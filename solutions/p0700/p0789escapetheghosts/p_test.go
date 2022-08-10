package p0789escapetheghosts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_escapeGhosts(t *testing.T) {
	for _, tc := range []struct {
		ghosts [][]int
		target []int
		want   bool
	}{
		{[][]int{{1, 0}, {0, 3}}, []int{0, 1}, true},
		{[][]int{{1, 0}}, []int{2, 0}, false},
		{[][]int{{2, 0}}, []int{1, 0}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ghosts), func(t *testing.T) {
			require.Equal(t, tc.want, escapeGhosts(tc.ghosts, tc.target))
		})
	}
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	manh := func(p1, p2 []int) int {
		return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	}
	dist := manh([]int{0, 0}, target)
	for _, g := range ghosts {
		if manh(g, target) <= dist {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package p1037validboomerang

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isBoomerang(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   bool
	}{
		{leetcode.ParseMatrix("[[52,86],[12,65],[24,71]]"), true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, isBoomerang(tc.points))
		})
	}
}

func isBoomerang(points [][]int) bool {
	p := points
	equal := func(p1, p2 []int) bool {
		return p1[0] == p2[0] && p1[1] == p2[1]
	}
	if equal(p[0], p[1]) || equal(p[1], p[2]) || equal(p[0], p[2]) {
		return false
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] < p[j][1]
		}
		return p[i][0] < p[j][0]
	})
	dx1 := float64(p[1][0] - p[0][0])
	dy1 := float64(p[1][1] - p[0][1])
	dx2 := float64(p[2][0] - p[1][0])
	dy2 := float64(p[2][1] - p[1][1])
	if dx1 == 0 || dx2 == 0 {
		return dx2 != dx1
	}
	if dy1 == 0 || dy2 == 0 {
		return dy2 != dy1
	}

	return dy1/dx1 != dy2/dx2
}

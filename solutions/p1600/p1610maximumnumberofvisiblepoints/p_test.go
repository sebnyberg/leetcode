package p1610maximumnumberofvisiblepoints

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_visiblePoints(t *testing.T) {
	for i, tc := range []struct {
		points   [][]int
		angle    int
		location []int
		want     int
	}{
		{
			leetcode.ParseMatrix("[[1,0],[2,1]]"),
			13, []int{1, 1}, 1,
		},
		{
			leetcode.ParseMatrix("[[2,1],[2,2],[3,3]]"),
			90, []int{1, 1}, 3,
		},
		{
			leetcode.ParseMatrix("[[2,1],[2,2],[3,4],[1,1]]"),
			90, []int{1, 1}, 4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, visiblePoints(tc.points, tc.angle, tc.location))
		})
	}
}

const eps = 1e-5

func visiblePoints(points [][]int, angle int, location []int) int {
	// It's clear what we have to do, but it's been a while since I did such a
	// mathsy problem.
	//
	// The idea is to calculate the angle for each point. In accordance with
	// standards, we denote the angle as the counter-clockwise angle from the
	// right-side x-axis.
	//
	angles := make([]float64, 0, len(points))
	x0 := float64(location[0])
	y0 := float64(location[1])
	var extra int
	for _, p := range points {
		x1 := float64(p[0])
		y1 := float64(p[1])
		dx := x1 - x0
		dy := y1 - y0

		if dx == 0 && dy == 0 {
			extra++
			continue
		}
		angle := (math.Atan2(dx, dy) / (2 * math.Pi)) * 360
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	n := len(angles)
	angles = append(angles, angles...)
	for i := n; i < 2*n; i++ {
		angles[i] += 360.0
	}
	var j int
	var res int
	for i := range angles {
		for angles[i]-angles[j] > float64(angle) {
			j++
		}
		res = max(res, i-j+1)
	}
	return extra + res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

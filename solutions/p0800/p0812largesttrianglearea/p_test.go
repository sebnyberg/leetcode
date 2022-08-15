package p0812largesttrianglearea

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_largestTriangleArea(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   float64
	}{
		// {
		// 	leetcode.ParseMatrix("[[0,0],[0,1],[1,0],[0,2],[2,0]]"),
		// 	2,
		// },
		{
			leetcode.ParseMatrix("[[35,-23],[-12,-48],[-34,-40],[21,-25],[-35,-44],[24,1],[16,-9],[41,4],[-36,-49],[42,-49],[-37,-20],[-35,11],[-2,-36],[18,21],[18,8],[-24,14],[-23,-11],[-8,44],[-19,-3],[0,-10],[-21,-4],[23,18],[20,11],[-42,24],[6,-19]]"),
			3627,
		},
		{
			leetcode.ParseMatrix("[[1,0],[0,0],[0,1]]"),
			0.5,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.InEpsilon(t, tc.want, largestTriangleArea(tc.points), 1e-5)
		})
	}
}

func largestTriangleArea(points [][]int) float64 {
	dist := func(p1, p2 []int) float64 {
		dx := p2[0] - p1[0]
		dy := p2[1] - p1[1]
		return math.Sqrt(float64(dx*dx + dy*dy))
	}
	calcArea := func(p1, p2, p3 []int) float64 {
		a := float64(dist(p1, p2))
		b := float64(dist(p1, p3))
		c := float64(dist(p2, p3))
		s := (a + b + c) / 2
		// Heron's formula
		return math.Sqrt(s * (s - a) * (s - b) * (s - c))
	}
	n := len(points)
	var maxArea float64
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				res := calcArea(points[i], points[j], points[k])
				if math.IsNaN(res) {
					continue
				}
				maxArea = math.Max(maxArea, res)
			}
		}
	}
	return maxArea
}

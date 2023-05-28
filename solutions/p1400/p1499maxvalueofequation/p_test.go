package p1499maxvalueofequation

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_findMaxValueOfEquation(t *testing.T) {
	for i, tc := range []struct {
		points [][]int
		k      int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[-19,-12],[-5,-18],[2,-2],[10,3],[11,-3],[13,17]]"), 13, 26,
		},
		{
			leetcode.ParseMatrix("[[1,3],[2,0],[5,10],[6,-10]]"), 1, 4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxValueOfEquation(tc.points, tc.k))
		})
	}
}

func findMaxValueOfEquation(points [][]int, k int) int {
	// Keep a queue in falling order of objectively "high" value
	q := [][]int{}
	res := math.MinInt64
	for _, p := range points {
		for len(q) > 0 && p[0]-q[0][0] > k {
			q = q[1:]
		}
		if len(q) > 0 {
			val := p[0] - q[0][0] + p[1] + q[0][1]
			res = max(res, val)
		}
		// Note: for the current point is objectively better than the previous
		// no matter what other point we are considering (given that it is in
		// range), it must compensate for the loss of x-distance in terms of
		// having a higher y value
		for len(q) > 0 {
			prev := q[len(q)-1]
			dy := p[1] - prev[1]
			dx := p[0] - prev[0]
			if dy >= dx {
				q = q[:len(q)-1]
			} else {
				break
			}
		}
		q = append(q, p)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

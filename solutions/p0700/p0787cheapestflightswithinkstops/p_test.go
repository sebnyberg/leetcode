package p0787cheapestflightswithinkstops

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_findCheapestPrice(t *testing.T) {
	for _, tc := range []struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{
			4,
			leetcode.ParseMatrix("[[0,1,1],[0,2,5],[1,2,1],[2,3,1]]"),
			0,
			3,
			1,
			6,
		},
		{
			4,
			leetcode.ParseMatrix("[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]"),
			0, 3, 1, 700,
		},
		{
			5,
			leetcode.ParseMatrix("[[4,1,1],[1,2,3],[0,3,2],[0,4,10],[3,1,1],[1,4,3]]"),
			2,
			1,
			1,
			-1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k))
		})
	}
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Perform BFS while <= k
	type edge struct{ to, weight int }
	edges := make([][]edge, n)
	for _, f := range flights {
		edges[f[0]] = append(edges[f[0]], edge{f[1], f[2]})
	}
	dists := make([]int, n)
	for i := range dists {
		dists[i] = math.MaxInt32
	}

	type visit struct{ to, dist int }
	curr := []visit{{src, 0}}
	next := []visit{}
	for kk := 0; kk <= k && len(curr) > 0; kk++ {
		next = next[:0]
		for _, x := range curr {
			for _, nei := range edges[x.to] {
				d := x.dist + nei.weight
				if dists[nei.to] > d {
					dists[nei.to] = d
					next = append(next, visit{nei.to, d})
				}
			}
		}
		curr, next = next, curr
	}
	if dists[dst] == math.MaxInt32 {
		return -1
	}

	return dists[dst]
}

package p1057campusbikes

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_assignBikes(t *testing.T) {
	for _, tc := range []struct {
		workers [][]int
		bikes   [][]int
		want    []int
	}{

		{
			leetcode.ParseMatrix("[[664,994],[3,425],[599,913],[220,352],[145,348],[604,428],[519,183],[732,148]]"),
			leetcode.ParseMatrix("[[611,698],[113,338],[579,770],[276,588],[948,679],[731,525],[925,877],[182,281],[39,299]]"),
			[]int{0, 8, 2, 7, 1, 5, 3, 4},
		},
		{
			leetcode.ParseMatrix("[[0,0],[2,1]]"),
			leetcode.ParseMatrix("[[1,2],[3,3]]"),
			[]int{1, 0},
		},
		{
			leetcode.ParseMatrix("[[0,0],[1,1],[2,0]]"),
			leetcode.ParseMatrix("[[1,0],[2,2],[2,1]]"),
			[]int{0, 2, 1},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.workers), func(t *testing.T) {
			require.Equal(t, tc.want, assignBikes(tc.workers, tc.bikes))
		})
	}
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	// Since there are 1000 elements, we can perform O(n^2) distance check for
	// all bikes and all workers. For something like 1e4 or 1e5, we'd need to do
	// a much more efficient check, probably discovering reasonable matches first
	// using binary search, then brute-force for the remainder.

	type match struct {
		worker int
		bike   int
		sqdist int
	}
	dist := func(p1, p2 []int) int {
		return abs(p2[0]-p1[0]) + abs(p2[1]-p1[1])
	}

	n := len(workers)
	matches := make([]match, 0, n*n)
	for i, w := range workers {
		for j, b := range bikes {
			matches = append(matches,
				match{
					worker: i,
					bike:   j,
					sqdist: dist(w, b),
				},
			)
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].sqdist == matches[j].sqdist {
			if matches[i].worker == matches[j].worker {
				return matches[i].bike < matches[j].bike
			}
			return matches[i].worker < matches[j].worker
		}
		return matches[i].sqdist < matches[j].sqdist
	})
	var bikeTaken [1001]bool
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	for _, m := range matches {
		if res[m.worker] != -1 || bikeTaken[m.bike] {
			continue
		}
		res[m.worker] = m.bike
		bikeTaken[m.bike] = true
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

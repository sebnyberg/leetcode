package p1615maximalnetworkrank

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maximalNetworkRank(t *testing.T) {
	for _, tc := range []struct {
		n     int
		roads [][]int
		want  int
	}{
		{4, leetcode.ParseMatrix("[[0,1],[0,3],[1,2],[1,3]]"), 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maximalNetworkRank(tc.n, tc.roads))
		})
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	// The only real issues are double counting a road between the pair of cities.
	//
	// Sort roads by starting node (and road pairs)
	for i := range roads {
		sort.Ints(roads[i])
	}
	sort.Slice(roads, func(i, j int) bool {
		if roads[i][0] == roads[j][0] {
			return roads[i][1] < roads[j][1]
		}
		return roads[i][0] < roads[j][0]
	})

	// Then count each pair to its city
	count := make([]int, n)
	for _, r := range roads {
		count[r[0]]++
		count[r[1]]++
	}

	// Then for each pair of cities, combine their count but deduct 1 if the two
	// are connected
	var res int
	var i int // current position in roads
	m := len(roads)
	for a := 0; a < n-1; a++ {
		for b := a + 1; b < n; b++ {
			pairCount := count[a] + count[b]
			if i < m && roads[i][0] == a && roads[i][1] == b {
				i++
				pairCount--
			}
			res = max(res, pairCount)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2225findplayerswithzerooronelosses

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWinners(t *testing.T) {
	for _, tc := range []struct {
		matches [][]int
		want    [][]int
	}{
		{
			leetcode.ParseMatrix("[[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]"),
			leetcode.ParseMatrix("[[1,2,10],[4,5,7,8]]"),
		},
		{
			leetcode.ParseMatrix("[[2,3],[1,3],[5,4],[6,4]]"),
			leetcode.ParseMatrix("[[1,2,5,6],[]]"),
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matches), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findWinners(tc.matches))
		})
	}
}

func findWinners(matches [][]int) [][]int {
	games := make([]int, 1e5+1)
	wins := make([]int, 1e5+1)
	for _, m := range matches {
		p1, p2 := m[0], m[1]
		games[p1]++
		games[p2]++
		wins[p1]++
	}
	winners := make([]int, 0, 1e5+1)
	lostOne := make([]int, 0, 1e5+1)
	for i := range games {
		if games[i] == 0 {
			continue
		}
		if games[i] == wins[i] {
			winners = append(winners, i)
		}
		if games[i]-wins[i] == 1 {
			lostOne = append(lostOne, i)
		}
	}
	return [][]int{winners, lostOne}
}

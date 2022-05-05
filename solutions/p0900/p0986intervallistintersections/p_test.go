package p0986intervallistintersections

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intervalIntersection(t *testing.T) {
	for _, tc := range []struct {
		firstList  [][]int
		secondList [][]int
		want       [][]int
	}{
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.firstList), func(t *testing.T) {
			require.Equal(t, tc.want, intervalIntersection(tc.firstList, tc.secondList))
		})
	}
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	const (
		intervalStart = false
		intervalEnd   = true
	)
	type intervalEdge struct {
		pos   int32
		delta bool
		// 60 bit padding
	}
	edges := make([]intervalEdge, 0, len(firstList)+len(secondList))
	lists := [][][]int{firstList, secondList}
	for _, list := range lists {
		for _, interval := range list {
			edges = append(edges, intervalEdge{
				pos:   int32(interval[0]),
				delta: intervalStart,
			})
			edges = append(edges, intervalEdge{
				pos:   int32(interval[1]),
				delta: intervalEnd,
			})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].pos == edges[j].pos {
			return !edges[i].delta
		}
		return edges[i].pos < edges[j].pos
	})
	var intervalCount int
	result := make([][]int, 0)
	for _, edge := range edges {
		// end of an interval
		if intervalCount == 2 {
			result[len(result)-1][1] = int(edge.pos)
		}
		if edge.delta == intervalStart {
			intervalCount++
		} else {
			intervalCount--
		}
		if intervalCount == 2 {
			result = append(result, []int{int(edge.pos), 0})
		}
	}
	return result
}

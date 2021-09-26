package p2015averageheightofbuildingsineachsegment

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_averageHeightOfBuildings(t *testing.T) {
	for _, tc := range []struct {
		buildings [][]int
		want      [][]int
	}{
		{[][]int{{1, 5, 2}, {3, 10, 4}}, [][]int{{1, 3, 2}, {3, 5, 3}, {5, 10, 4}}},
		{[][]int{{1, 3, 2}, {2, 5, 3}, {2, 8, 3}}, [][]int{{1, 3, 2}, {3, 8, 3}}},
		{[][]int{{1, 2, 1}, {5, 6, 1}}, [][]int{{1, 2, 1}, {5, 6, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buildings), func(t *testing.T) {
			require.Equal(t, tc.want, averageHeightOfBuildings(tc.buildings))
		})
	}
}

func averageHeightOfBuildings(buildings [][]int) [][]int {
	deltas := make([][]int, 0, len(buildings)+1)
	for _, b := range buildings {
		deltas = append(deltas, []int{b[0], b[2]})
		deltas = append(deltas, []int{b[1], -b[2]})
	}
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i][0] < deltas[j][0]
	})
	// Compact deltas
	positions := []int{deltas[0][0]}
	deltasForPos := [][]int{{deltas[0][1]}}
	for i, posIdx := 1, 0; i < len(deltas); i++ {
		pos, heightDelta := deltas[i][0], deltas[i][1]
		if pos == positions[posIdx] {
			deltasForPos[posIdx] = append(deltasForPos[posIdx], heightDelta)
		} else {
			positions = append(positions, pos)
			deltasForPos = append(deltasForPos, []int{heightDelta})
			posIdx++
		}
	}

	// Evaluate each position in the list of deltas and construct the result
	start := positions[0]
	count := len(deltasForPos[0])
	var curHeight int
	for _, h := range deltasForPos[0] {
		curHeight += h
	}
	avg := curHeight / count
	res := make([][]int, 0)
	// For each position in deltas, calculate the new average.
	// If the average changed, push to result list.
	for posIdx := 1; posIdx < len(positions); posIdx++ {
		for _, h := range deltasForPos[posIdx] {
			curHeight += h
			if h < 0 {
				count--
			} else {
				count++
			}
		}
		if count == 0 || posIdx == len(positions)-1 {
			res = append(res, []int{start, positions[posIdx], avg})
			start = -1
			avg = 0
			continue
		}
		newAvg := curHeight / count
		if newAvg != avg {
			if start != -1 {
				res = append(res, []int{start, positions[posIdx], avg})
			}
			start = positions[posIdx]
		}
		avg = newAvg
	}
	return res
}

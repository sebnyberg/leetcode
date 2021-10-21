package p1996thenumberofweakcharactersinthegame

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWeakCharacters(t *testing.T) {
	for _, tc := range []struct {
		properties [][]int
		want       int
	}{
		{[][]int{{5, 5}, {6, 3}, {3, 6}}, 0},
		{[][]int{{2, 2}, {3, 3}}, 1},
		{[][]int{{1, 5}, {10, 4}, {4, 3}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.properties), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWeakCharacters(tc.properties))
		})
	}
}

func numberOfWeakCharacters(properties [][]int) int {
	// Sort by attack
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0]
	})
	// Bin by attack
	maxDef := make([]int, 1)
	maxDef[0] = properties[0][1]
	defForAttack := make([][]int, 1)
	defForAttack[0] = []int{properties[0][1]}
	var binIdx int
	binVals := make([]int, 1)
	binVals[0] = properties[0][0]
	for _, prop := range properties[1:] {
		if prop[0] > binVals[binIdx] {
			binIdx++
			defForAttack = append(defForAttack, []int{prop[1]})
			binVals = append(binVals, prop[0])
			maxDef = append(maxDef, prop[1])
			continue
		}
		maxDef[binIdx] = max(maxDef[binIdx], prop[1])
		defForAttack[binIdx] = append(defForAttack[binIdx], prop[1])
	}
	maxDefAbove := make([]int, len(maxDef))
	for i := len(maxDef) - 2; i >= 0; i-- {
		maxDefAbove[i] = max(maxDefAbove[i+1], maxDef[i+1])
	}
	var res int
	for i := 0; i < len(binVals)-1; i++ {
		for _, def := range defForAttack[i] {
			if def < maxDefAbove[i] {
				res++
			}
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

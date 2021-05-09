package p1854maximumpopulationyear

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumPopulation(t *testing.T) {
	for _, tc := range []struct {
		logs [][]int
		want int
	}{
		{[][]int{{1993, 1999}, {2000, 2010}}, 1993},
	} {
		t.Run(fmt.Sprintf("%+v", tc.logs), func(t *testing.T) {
			require.Equal(t, tc.want, maximumPopulation(tc.logs))
		})
	}
}

func maximumPopulation(logs [][]int) int {
	var alive [100]int
	for _, log := range logs {
		birthIdx, deathIdx := log[0]-1950, log[1]-1950
		for i := birthIdx; i < deathIdx; i++ {
			alive[i]++
		}
	}
	var maxAlive int
	var maxYear int
	for i, aliveForYear := range alive {
		if aliveForYear > maxAlive {
			maxAlive = aliveForYear
			maxYear = i + 1950
		}
	}
	return maxYear
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1029twocityscheduling

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoCitySchedCost(t *testing.T) {
	for _, tc := range []struct {
		costs [][]int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[10,20],[30,200],[400,50],[30,20]]"),
			110,
		},
		{
			leetcode.ParseMatrix("[[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]"),
			1859,
		},
		{
			leetcode.ParseMatrix("[[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]"),
			3086,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.costs), func(t *testing.T) {
			require.Equal(t, tc.want, twoCitySchedCost(tc.costs))
		})
	}
}

func twoCitySchedCost(costs [][]int) int {
	// The cost of going to each city is already available, so no path-finding is
	// required. This is rather a matching problem.

	// For each pair of cities, there is a gain / loss from putting a person in
	// one or the other city. If we subtract the cost on one side with the cost of
	// the other, we get the gain / loss for each person. We should aim to resolve
	// the highest deltas first.
	m := len(costs[0])
	n := len(costs) / 2

	type delta struct {
		idx int
		val int
	}
	deltas := make([]delta, n*2)
	minSum := math.MaxInt32
	for i := 0; i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			// Calculate deltas for each pair of cities / people
			for k := range costs {
				deltas[k] = delta{idx: k, val: costs[k][j] - costs[k][i]}
			}
			sort.Slice(deltas, func(i, j int) bool {
				return abs(deltas[i].val) > abs(deltas[j].val)
			})

			var sum int
			var count [2]int
			for _, d := range deltas {
				if count[0] < n && (d.val >= 0 || count[1] == n) { // Second has larger cost, pick first
					sum += costs[d.idx][i]
					count[0]++
				} else { // First has larger cost, pick second
					sum += costs[d.idx][j]
					count[1]++
				}
			}
			minSum = min(minSum, sum)
		}
	}
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

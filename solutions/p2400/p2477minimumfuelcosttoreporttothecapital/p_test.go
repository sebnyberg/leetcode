package p2477minimumfuelcosttoreporttothecapital

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumFuelCost(t *testing.T) {
	for i, tc := range []struct {
		roads [][]int
		seats int
		want  int64
	}{
		{
			leetcode.ParseMatrix("[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]]"),
			2, 7,
		},
		{
			leetcode.ParseMatrix("[[0,1],[0,2],[0,3]]"),
			5, 3,
		},
		{
			leetcode.ParseMatrix("[[]]"),
			1, 0,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumFuelCost(tc.roads, tc.seats))
		})
	}
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	// For each node, calculate the number of people leaving that node. The
	// number of cars needed to accomodate those people is that node's
	// contribution to the result
	if len(roads) == 0 {
		return 0
	}
	var n int
	adj := make([][]int, 0)
	for _, x := range roads {
		for n <= x[0] {
			n++
			adj = append(adj, []int{})
		}
		for n <= x[1] {
			n++
			adj = append(adj, []int{})
		}
		adj[x[0]] = append(adj[x[0]], x[1])
		adj[x[1]] = append(adj[x[1]], x[0])
	}
	seen := make([]bool, n)
	seen[0] = true
	var total int64
	for _, j := range adj[0] {
		seen[j] = true
		peopleLeavingNode(seen, adj, seats, j, &total)
	}
	return total
}

func peopleLeavingNode(seen []bool, adj [][]int, seats, i int, total *int64) int64 {
	var npeople int64 = 1
	for _, j := range adj[i] {
		if seen[j] {
			continue
		}
		seen[j] = true
		npeople += peopleLeavingNode(seen, adj, seats, j, total)
	}
	a := ((npeople - 1) / int64(seats)) + 1
	*total += a
	return npeople
}

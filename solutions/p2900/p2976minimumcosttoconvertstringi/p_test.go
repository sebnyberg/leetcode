package p2976minimumcosttoconvertstringi

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCost(t *testing.T) {
	for _, tc := range []struct {
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
		want     int
	}{
		{"abcd", "acbe", []byte("abcced"), []byte("bcbebe"), []int{2, 5, 5, 1, 2, 20}, 28},
	} {
		t.Run(fmt.Sprintf("%+v", tc.source), func(t *testing.T) {
			require.Equal(t, tc.want, minimumCost(tc.source, tc.target, tc.original, tc.changed, tc.cost))
		})
	}
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	// The choice of change is not unique, so we have to use some kind of
	// traversal algorithm to determine the minimum cost.
	//
	// Two options are bellman-ford and dijstra. We can use bellman-ford.

	var minCost [26][26]int
	for i := range minCost {
		for j := range minCost {
			minCost[i][j] = math.MaxInt32
		}
		minCost[i][i] = 0
	}

	type edge struct {
		from   byte
		to     byte
		weight int
	}
	var adj [26][]edge
	for i := range original {
		a := original[i] - 'a'
		b := changed[i] - 'a'
		adj[a] = append(adj[a], edge{from: a, to: b, weight: cost[i]})
	}

	var q []byte
	traverse := func(start byte) {
		q = append(q[:0], start)
		for i := 0; i < len(q); i++ {
			x := q[i]
			if x >= 26 {
				panic(fmt.Sprintf("%v, %v\n", start, x))
			}
			for _, e := range adj[x] {
				if minCost[start][e.to] <= minCost[start][x]+e.weight {
					continue
				}
				minCost[start][e.to] = minCost[start][x] + e.weight
				q = append(q, byte(e.to))
			}
		}
	}
	for i := range minCost {
		traverse(byte(i))
	}
	var res int
	for i := range source {
		c := minCost[source[i]-'a'][target[i]-'a']
		if c == math.MaxInt32 {
			return -1
		}
		res += minCost[source[i]-'a'][target[i]-'a']
	}
	return int64(res)
}

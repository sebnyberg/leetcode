package p1595minimumcosttoconnecttrogroupsofpoints

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_connectTwoGroups(t *testing.T) {
	for i, tc := range []struct {
		cost [][]int
		want int
	}{
		{[][]int{{15, 96}, {36, 2}}, 17},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, connectTwoGroups(tc.cost))
		})
	}
}

type state struct {
	i  int
	bm int
}

func connectTwoGroups(cost [][]int) int {
	// Let's find a way to describe the current "state of the world" in order to
	// use DP.
	//
	// We could describe the state in terms of having used a connection or not.
	// That would give us O(n^n) = insane state space.
	//
	// We could also describe the state in terms of a node being covered or not,
	// which gives O(2^(2*n)) states, but many of these states are impossible.
	// For example, a node cannot be assigned without another node on the other
	// side being assigned as well.
	//
	// A worst-case of 2^24 ~= 8M - it should be OK.
	//
	mem := make(map[state]int)
	res := dp(mem, cost, 0, 0)
	return res
}

func dp(mem map[state]int, cost [][]int, i, bm int) int {
	key := state{i, bm}
	if v, exists := mem[key]; exists {
		return v
	}
	if i == len(cost) {
		// Reached right-hand side
		// Any not connected node should choose its optimal edge.
		var res int
		for j := 0; j < len(cost[0]); j++ {
			if bm&(1<<j) > 0 {
				continue
			}
			// Find the optimal edge
			minCost := math.MaxInt32
			for i := range cost {
				minCost = min(minCost, cost[i][j])
			}
			res += minCost
		}
		mem[key] = res
		return res
	}

	// Try connecting this node to any other node on the right side
	res := math.MaxInt32
	for j := 0; j < len(cost[0]); j++ {
		res = min(res, cost[i][j]+dp(mem, cost, i+1, bm|(1<<j)))
	}
	mem[key] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

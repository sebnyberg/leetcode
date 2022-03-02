package p0514freedomtrail

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRotateSteps(t *testing.T) {
	for _, tc := range []struct {
		ring string
		key  string
		want int
	}{
		{"godding", "gd", 4},
		{"godding", "godding", 13},
		{
			"caotmcaataijjxi",
			"oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx",
			137,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ring), func(t *testing.T) {
			require.Equal(t, tc.want, findRotateSteps(tc.ring, tc.key))
		})
	}
}

func findRotateSteps(ring string, key string) int {
	// The problem with this exercise is that the closest possible letter may be
	// the best choice for a single letter, but it's not necessarily the best
	// choice for the whole sequence.

	// One way to try to find the best solution is to consider each letter to be
	// a node in a graph. Then if there are two ways to go to find the closest
	// letter, try both alternatives. Maybe we need to even try all alternatives?

	// Dijkstra's took too much memory and caused a runtime error.

	// Let's instead try DFS - it's less memory intensive.
	// Is there any invariant that we can use to reduce pressure?
	// Well.. if we have arrived at the same index at a key index higher than
	// or equal to the current and with a lower or equal distance, then there
	// is no reason to continue

	var charIndices [26][]int
	for i, ch := range ring {
		charIndices[ch-'a'] = append(charIndices[ch-'a'], i)
	}

	mem := make(map[state]int, len(key))
	n := len(ring)
	m := len(key)
	res := dfs(mem, &charIndices, ring, key, 0, 0, n, m)

	return res
}

func dfs(
	mem map[state]int,
	charIndices *[26][]int,
	ring, key string,
	idx, keyIdx, n, m int,
) int {
	if keyIdx == m {
		return 0
	}
	k := state{idx, keyIdx}
	if v, exists := mem[k]; exists {
		return v
	}
	res := math.MaxInt32
	for _, nearIdx := range charIndices[key[keyIdx]-'a'] {
		dist := min(abs(idx-nearIdx), abs(n+idx-nearIdx))
		dist = min(dist, abs(n+nearIdx-idx))
		res = min(res, dist+dfs(mem, charIndices, ring, key, nearIdx, keyIdx+1, n, m))
	}
	mem[k] = res + 1
	return mem[k]
}

type state struct {
	keyIdx int
	idx    int
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

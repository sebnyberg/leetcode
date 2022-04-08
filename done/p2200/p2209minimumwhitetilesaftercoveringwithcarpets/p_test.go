package p2209

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumWhiteTiles(t *testing.T) {
	for _, tc := range []struct {
		floor      string
		numCarpets int
		carpetLen  int
		want       int
	}{
		{"1110111", 2, 1, 4},
		{"01101", 1, 2, 1},
		{"101101101", 2, 2, 2},
		{"0000", 1, 1, 0},
		{"11111", 2, 3, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.floor), func(t *testing.T) {
			require.Equal(t, tc.want, minimumWhiteTiles(tc.floor, tc.numCarpets, tc.carpetLen))
		})
	}
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	// We might need to consider each possible position for each carpet.
	n := len(floor)

	right := make([]int, n)
	var alreadyCovered int
	for i := n - 1; i >= 0; i-- {
		if floor[i] == '0' {
			right[i]++
			alreadyCovered++
		}
		if i < n-1 {
			right[i] += right[i+1]
		}
	}

	// For each position, you should either place the carpet or not.
	mem := make(map[key]int)
	a := dfs(mem, floor, right, 0, numCarpets, carpetLen, n)
	res := n - alreadyCovered - a
	// res := coveredSoFar + dfs(mem, right, 0, numCarpets, carpetLen, n)
	return res
}

type key struct {
	idx        int
	numCarpets int
}

func dfs(mem map[key]int, floor string, right []int, i, carpetRemains, carpetLen, n int) int {
	k := key{i, carpetRemains}
	if v, exists := mem[k]; exists {
		return v
	}
	if i >= n {
		return 0
	}
	if carpetRemains == 0 {
		return 0
	}

	// Either don't place mat in this position, i.e. skip
	res := dfs(mem, floor, right, i+1, carpetRemains, carpetLen, n)

	// Or place the mat
	var placeRes int
	for j := i; j < min(n, i+carpetLen); j++ {
		if floor[j] == '1' {
			placeRes++
		}
	}
	res = max(res, placeRes+dfs(mem, floor, right, i+carpetLen, carpetRemains-1, carpetLen, n))
	mem[k] = res
	return mem[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

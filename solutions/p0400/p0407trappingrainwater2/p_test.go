package p0407trappingrainwater2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"

	"github.com/stretchr/testify/require"
)

func Test_trapRainWater(t *testing.T) {
	for _, tc := range []struct {
		heightMap [][]int
		want      int
	}{
		{leetcode.ParseMatrix("[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]"), 4},
		{leetcode.ParseMatrix("[[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]"), 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.heightMap), func(t *testing.T) {
			require.Equal(t, tc.want, trapRainWater(tc.heightMap))
		})
	}
}

func trapRainWater(heightMap [][]int) int {
	// Everything is a "puddle" until it is joined toward an edge of the map.
	// We can join locations together as a DSU and calculate the total trapped
	// rainwater whenever it is joned toward the edge.

	m := len(heightMap)
	n := len(heightMap[0])
	parent := make([]int, m*n)
	trapped := make([]bool, m*n)
	for i := range parent {
		parent[i] = i
		trapped[i] = true
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		ra := find(parent[a])
		parent[a] = ra // path compression
		return ra
	}

	// Union joins together two puddles. If one puddle is trapped and the other is
	// not, then we add the trapped water to the total result and mark it as
	// untrapped.
	union := func(a, b, currentHeight int) int {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return 0
		}
		var trappedWater int
		if !trapped[rb] {
			ra, rb = rb, ra
		}
		if !trapped[ra] && trapped[rb] {
			// calculate total amount trapped prior to this puddle going over the
			// edge of the map.
			for i := range heightMap {
				for j, h := range heightMap[i] {
					if find(i*n+j) == rb {
						trappedWater += currentHeight - h
					}
				}
			}
		}
		parent[rb] = ra
		return trappedWater
	}

	// Mark edge blocks as untrapped
	for i := range heightMap {
		trapped[i*n] = false
		trapped[(i+1)*n-1] = false
	}
	for j := range heightMap[0] {
		trapped[j] = false
		trapped[n*(m-1)+j] = false
	}

	// Partition blocks by height
	unitsByHeight := make(map[int][][2]int)
	for i := range heightMap {
		for j, v := range heightMap[i] {
			unitsByHeight[v] = append(unitsByHeight[v], [2]int{i, j})
		}
	}
	var heights []int
	for k := range unitsByHeight {
		heights = append(heights, k)
	}
	sort.Ints(heights)

	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Iterate over blocks for each height.
	var totalWater int
	for _, h := range heights {
		for _, block := range unitsByHeight[h] {
			i := block[0]
			j := block[1]
			for _, d := range dirs {
				ii := i + d[0]
				jj := j + d[1]
				if !ok(ii, jj) || heightMap[ii][jj] > h {
					continue
				}
				// The neighbouring block is at a lower level, join it with this puddle.
				totalWater += union(i*n+j, ii*n+jj, h)
			}
		}
	}

	return totalWater
}

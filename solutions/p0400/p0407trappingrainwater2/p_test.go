package p0407trappingrainwater2

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

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
	m, n := len(heightMap), len(heightMap[0])

	// Find each distinct level in the elevation map
	levels := make([]int, 0, m*n)
	for i := range heightMap {
		for _, v := range heightMap[i] {
			levels = append(levels, v)
		}
	}
	sort.Ints(levels)
	var j int
	for i := range levels {
		if i > 0 && levels[i] == levels[i-1] {
			continue
		}
		levels[j] = levels[i]
		j++
	}
	levels = levels[:j]

	// Helper functions
	dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}
	isEdge := func(i, j int) bool {
		return i == 0 || i == m-1 || j == 0 || j == n-1
	}

	// fill starts in (i, j) and flood-fills the area with the value v
	// If this results in filling an edge position, the return value is zero
	fill := func(i, j, val int) int {
		start := heightMap[i][j]
		if start == val {
			return 0
		}
		cur := [][]int{{i, j}}
		next := [][]int{}
		count := 1
		heightMap[i][j] = val
		var edge bool
		for len(cur) > 0 {
			next = next[:0]
			for _, p := range cur {
				if isEdge(p[0], p[1]) {
					edge = true
				}
				for _, dir := range dirs {
					ii := p[0] + dir[0]
					jj := p[1] + dir[1]
					if !ok(ii, jj) || heightMap[ii][jj] != start {
						continue
					}
					heightMap[ii][jj] = val
					count++
					next = append(next, []int{ii, jj})
				}
			}
			cur, next = next, cur
		}
		if edge {
			return 0
		}
		return count
	}

	// For each distinct level (except the last)
	var result int
	for k := 0; k < len(levels)-1; k++ {
		l := levels[k]
		delta := levels[k+1] - levels[k]
		// For each place in the height map that can be filled with water
		for i := range heightMap {
			for j, v := range heightMap[i] {
				if v != l {
					continue
				}
				// Fill with the next level and add the filled area to the result
				result += fill(i, j, levels[k+1]) * delta
			}
		}
	}
	return result
}

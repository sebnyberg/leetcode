package p1857largestcolorvalueindirectedgraph

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestPathValue(t *testing.T) {
	for _, tc := range []struct {
		colors string
		edges  [][]int
		want   int
	}{
		{"hhqhuqhqff",
			[][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}},
			3,
		},
		{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.colors), func(t *testing.T) {
			require.Equal(t, tc.want, largestPathValue(tc.colors, tc.edges))
		})
	}
}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	indeg := make([]int, n)
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		indeg[b]++
		adj[a] = append(adj[a], b)
	}
	q := list.New()
	seen := 0
	colorCounts := make([][26]int, n)
	for i, deg := range indeg {
		if deg == 0 {
			q.PushBack(i)
			colorCounts[i][colors[i]-'a'] = 1
		}
	}

	var maxColors int
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		maxColors = max(maxColors, maxInts(colorCounts[u]))
		seen++
		for _, nei := range adj[u] {
			indeg[nei]--
			neiColor := int(colors[nei] - 'a')
			for c := 0; c < 26; c++ {
				if c == neiColor {
					colorCounts[nei][c] = max(colorCounts[nei][c], colorCounts[u][c]+1)
				} else {
					colorCounts[nei][c] = max(colorCounts[nei][c], colorCounts[u][c])
				}
			}
			if indeg[nei] == 0 {
				q.PushBack(nei)
			}
		}
	}
	if seen < n {
		return -1
	}
	return maxColors
}

func maxInts(nums [26]int) int {
	var max int
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

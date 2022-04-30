package p0785graphbipartite2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isBipartite(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  bool
	}{
		{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{[][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}, false},
		{[][]int{{4}, {}, {4}, {4}, {0, 2, 3}}, true},
		{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, isBipartite(tc.graph))
		})
	}
}

type color = uint8

const (
	colorRed  = 0
	colorBlue = 1
)

func isBipartite(graph [][]int) bool {
	// Perform graph coloring.
	n := len(graph)
	visited := make([]bool, n)
	colors := make([]color, n)
	curr := []int{}
	next := []int{}
	var c color
	for i := range graph {
		if visited[i] {
			continue
		}
		curr = curr[:0]
		c = colorRed
		curr = append(curr, i)
		for len(curr) > 0 {
			next = next[:0]
			for _, j := range curr {
				if visited[j] {
					if colors[j] != c {
						return false
					}
					continue
				}
				visited[j] = true
				colors[j] = c
				next = append(next, graph[j]...)
			}
			curr, next = next, curr
			c = 1 - c
		}
	}
	return true
}

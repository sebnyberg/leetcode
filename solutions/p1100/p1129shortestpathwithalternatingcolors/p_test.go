package p1129shortestpathwithalternatingcolors

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	for i, tc := range []struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
		want      []int
	}{
		{3, [][]int{{0, 1}}, [][]int{{2, 1}}, []int{0, 1, -1}},
		{3, [][]int{{0, 1}, {1, 2}}, [][]int{}, []int{0, 1, -1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, shortestAlternatingPaths(tc.n, tc.redEdges, tc.blueEdges))
		})
	}
}

const red = 0
const blue = 1

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = math.MaxInt32
	}

	// Parse adjacency lists
	var adj [2][][]int
	adj[red] = make([][]int, n)
	adj[blue] = make([][]int, n)
	for _, e := range redEdges {
		a := e[0]
		b := e[1]
		adj[red][a] = append(adj[red][a], b)
	}
	for _, e := range blueEdges {
		a := e[0]
		b := e[1]
		adj[blue][a] = append(adj[blue][a], b)
	}

	// Perform bfs, avoiding visiting the same node twice.
	seen := make([][2]bool, n)
	curr := []int{}
	next := []int{}

	bfs := func(color int) {
		for i := range seen {
			seen[i] = [2]bool{false, false}
		}
		seen[0][color] = true
		curr = append(curr, 0)
		for k := 1; len(curr) > 0; k++ {
			next = next[:0]
			for _, x := range curr {
				for _, v := range adj[color][x] {
					if seen[v][1-color] {
						continue
					}
					if res[v] > k {
						res[v] = k
					}
					seen[v][1-color] = true
					next = append(next, v)
				}
			}
			color = 1 - color
			curr, next = next, curr
		}
	}

	bfs(blue)
	bfs(red)
	res[0] = 0
	for i := range res {
		if res[i] == math.MaxInt32 {
			res[i] = -1
		}
	}
	return res
}

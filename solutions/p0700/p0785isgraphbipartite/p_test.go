package p0785isgraphbipartite

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_isBipartite(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  bool
	}{
		{
			leetcode.ParseMatrix("[[1,2,3],[0,2],[0,1,3],[0,2]]"),
			false,
		},
		{
			leetcode.ParseMatrix("[[1,3],[0,2],[1,3],[0,2]]"),
			true,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, isBipartite(tc.graph))
		})
	}
}

func isBipartite(graph [][]int) bool {
	adj := graph
	n := len(adj)
	color := make([]uint8, n)
	curr := []int{0}
	next := []int{0}
	const red = 1
	const blue = 2
	for i := 0; i < n; i++ {
		if color[i] != 0 {
			continue
		}
		// Found a node to start BFS from
		curr = curr[:1]
		curr[0] = i
		color[i] = red
		var nextColor uint8 = blue
		for len(curr) > 0 {
			next = next[:0]
			for _, idx := range curr {
				for _, nei := range adj[idx] {
					if color[nei] == 0 {
						color[nei] = nextColor
						next = append(next, nei)
					} else {
						if color[nei] != nextColor {
							return false
						}
					}
				}
			}

			nextColor = (1 + (1 - (nextColor - 1)))
			curr, next = next, curr
		}
	}
	return true
}

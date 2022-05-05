package p0675cutofftreesforgolfevent

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_cutOffTree(t *testing.T) {
	for _, tc := range []struct {
		forest [][]int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[54581641,64080174,24346381,69107959],[86374198,61363882,68783324,79706116],[668150,92178815,89819108,94701471],[83920491,22724204,46281641,47531096],[89078499,18904913,25462145,60813308]]"),
			57,
		},
		{
			leetcode.ParseMatrix("[[1,2,3],[0,0,4],[7,6,5]]"),
			6,
		},
		{
			leetcode.ParseMatrix("[[1,2,3],[0,0,0],[7,6,5]]"),
			-1,
		},
		{
			leetcode.ParseMatrix("[[2,3,4],[0,0,5],[8,7,6]]"),
			6,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.forest), func(t *testing.T) {
			require.Equal(t, tc.want, cutOffTree(tc.forest))
		})
	}
}

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && forest[i][j] >= 1
	}
	if forest[0][0] < 1 {
		return -1
	}

	type treePos struct {
		height int
		pos    [2]int
	}
	trees := make([]treePos, 0, 1)
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] > 1 {
				trees = append(trees, treePos{
					height: forest[i][j],
					pos:    [2]int{i, j},
				})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].height < trees[j].height
	})

	// Since its a quite small grid, it should be possible to BFS from each
	// position and stop when the shortest path has been found.
	var totalSteps int
	startPos := [2]int{0, 0}
	for i := 0; i < len(trees); i++ {
		to := trees[i]
		if to.pos == startPos {
			continue
		}

		steps := 1
		curr := [][2]int{startPos}
		next := [][2]int{}

		var seen [51][51]bool
		seen[startPos[0]][startPos[1]] = true
		for len(curr) > 0 {
			next = next[:0]
			for _, p := range curr {
				for _, d := range dirs {
					ii, jj := p[0]+d[0], p[1]+d[1]
					if !ok(ii, jj) || seen[ii][jj] {
						continue
					}
					seen[ii][jj] = true
					if [2]int{ii, jj} == to.pos {
						totalSteps += steps
						goto continueWalk
					}
					next = append(next, [2]int{ii, jj})
				}
			}
			curr, next = next, curr
			steps++
		}
		return -1

	continueWalk:
		startPos = to.pos
	}

	return totalSteps
}

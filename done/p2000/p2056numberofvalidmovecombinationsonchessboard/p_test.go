package p2056numberofvalidmovecombinationsonchessboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countCombinations(t *testing.T) {
	for _, tc := range []struct {
		pieces    []string
		positions [][]int
		want      int
	}{
		{[]string{"rook", "rook"}, [][]int{{1, 1}, {8, 8}}, 223},
		{[]string{"rook"}, [][]int{{1, 1}}, 15},
		{[]string{"queen"}, [][]int{{1, 1}}, 22},
		{[]string{"bishop"}, [][]int{{4, 3}}, 12},
		{[]string{"queen", "bishop"}, [][]int{{5, 7}, {3, 4}}, 281},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pieces), func(t *testing.T) {
			require.Equal(t, tc.want, countCombinations(tc.pieces, tc.positions))
		})
	}
}

var pieceDirs = map[string][][]int8{
	"rook":   {{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
	"queen":  {{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}},
	"bishop": {{1, -1}, {1, 1}, {-1, 1}, {-1, -1}},
}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	var f combFinder
	var initialPos [4][2]int8
	for i := range positions {
		initialPos[i][0] = int8(positions[i][0]) - 1
		initialPos[i][1] = int8(positions[i][1]) - 1
	}
	f.combs = make(map[[4][2]int8]struct{})
	f.pickDirs(0, n, pieces, make([][]int8, n), &initialPos)
	return len(f.combs) + 1
}

type combFinder struct {
	combs map[[4][2]int8]struct{}
}

func (f *combFinder) pickDirs(i, n int, pieces []string, dirs [][]int8, initialPos *[4][2]int8) {
	if i == n {
		for movingMask := 1; movingMask < (1 << n); movingMask++ {
			f.findComb(*initialPos, dirs, movingMask, n)
		}
		return
	}
	for _, dir := range pieceDirs[pieces[i]] {
		dirs[i] = dir
		f.pickDirs(i+1, n, pieces, dirs, initialPos)
	}
}

func (f *combFinder) findComb(pos [4][2]int8, dirs [][]int8, movingMask, n int) {
	var seen [8][8]bool
	for i := 0; i < n; i++ {
		pos[i][0] += int8((movingMask>>i)&1) * dirs[i][0]
		pos[i][1] += int8((movingMask>>i)&1) * dirs[i][1]
		if pos[i][0] < 0 || pos[i][0] >= 8 || pos[i][1] < 0 || pos[i][1] >= 8 ||
			seen[pos[i][0]][pos[i][1]] {
			return
		}
		seen[pos[i][0]][pos[i][1]] = true
	}

	f.combs[pos] = struct{}{}

	// For each possible decision to move / stop moving a piece, continue search
	for nextMovingMask := 1; nextMovingMask < (1 << n); nextMovingMask++ {
		if nextMovingMask&movingMask != nextMovingMask {
			continue
		}
		f.findComb(pos, dirs, nextMovingMask, n)
	}
}

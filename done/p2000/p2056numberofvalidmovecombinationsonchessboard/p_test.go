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

var pieceDirs = map[string][][]int{
	"rook":   {{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
	"queen":  {{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}},
	"bishop": {{1, -1}, {1, 1}, {-1, 1}, {-1, -1}},
}

func countCombinations(pieces []string, positions [][]int) int {
	// Brute-force
	n := len(pieces)
	var f combFinder
	var pos [4][2]int8
	for i := range positions {
		pos[i][0] = int8(positions[i][0]) - 1
		pos[i][1] = int8(positions[i][1]) - 1
	}
	f.combs = make(map[[4][2]int8]struct{})
	f.pickDirs(0, n, pieces, make([][]int, n), &pos)
	return len(f.combs)
}

type combFinder struct {
	combs map[[4][2]int8]struct{}
}

func (f *combFinder) pickDirs(i, n int, pieces []string, dirs [][]int, pos *[4][2]int8) {
	if i == n {
		canMove := (1 << n) - 1
		f.findComb(*pos, dirs, canMove, n)
		return
	}
	for _, dir := range pieceDirs[pieces[i]] {
		dirs[i] = dir
		f.pickDirs(i+1, n, pieces, dirs, pos)
	}
}

func (f *combFinder) findComb(pos [4][2]int8, dirs [][]int, canMove int, n int) {
	f.combs[pos] = struct{}{}
	if canMove == 0 {
		return
	}
	move := func(toMove int) {
		// Move pieces
		for i := 0; i < n; i++ {
			if toMove&(1<<i) == 0 { // skip stopped pieces
				continue
			}
			pos[i][0] += int8(dirs[i][0])
			pos[i][1] += int8(dirs[i][1])
			defer func(j int) {
				pos[j][0] -= int8(dirs[j][0])
				pos[j][1] -= int8(dirs[j][1])
			}(i)
			if pos[i][0] < 0 || pos[i][0] >= 8 || pos[i][1] < 0 || pos[i][1] >= 8 {
				return
			}
		}
		// Check that the number of unique positions is equal to the number of pieces
		uniquePos := make(map[[2]int8]struct{})
		for i := 0; i < n; i++ {
			uniquePos[pos[i]] = struct{}{}
		}
		if len(uniquePos) != n {
			return
		}
		f.findComb(pos, dirs, toMove, n)
	}

	// For each possible decision to move / stop a piece, continue searching
	for toMove := 1; toMove < (1 << n); toMove++ {
		if toMove&canMove != toMove {
			continue
		}
		move(toMove)
	}
}

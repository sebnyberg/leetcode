package p2061numberofspacescleaningrobotcleaned

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfCleanRooms(t *testing.T) {
	for _, tc := range []struct {
		room [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}, 7},
		{[][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.room), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfCleanRooms(tc.room))
		})
	}
}

const (
	right = iota
	down
	left
	up
)

type position struct{ i, j int16 }

var delta = [4][2]int16{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func (p position) move(dir int) position {
	return position{
		i: p.i + delta[dir][0],
		j: p.j + delta[dir][1],
	}
}

func numberOfCleanRooms(room [][]int) int {
	m, n := int16(len(room)), int16(len(room[0]))
	// ok checks whether a position is valid
	ok := func(p position) bool {
		return p.i >= 0 && p.i < m &&
			p.j >= 0 && p.j < n &&
			room[p.i][p.j] != 1
	}
	pos := position{0, 0}
	dir := right
	var count int
	for {
		if room[pos.i][pos.j] == 0 {
			count++
		}
		if room[pos.i][pos.j]&(1<<(dir+1)) > 0 { // loop
			return count
		}
		room[pos.i][pos.j] |= 1 << (dir + 1)
		for i := 0; !ok(pos.move(dir)); i++ {
			if i == 4 { // deadlock
				return count
			}
			dir = (dir + 1) % 4
		}
		pos = pos.move(dir)
	}
}

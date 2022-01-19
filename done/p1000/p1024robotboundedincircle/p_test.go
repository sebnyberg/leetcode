package p1024robotboundedincircle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isRobotBounded(t *testing.T) {
	for _, tc := range []struct {
		instructions string
		want         bool
	}{
		{"GGLLGG", true},
		{"GG", false},
		{"GL", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.instructions), func(t *testing.T) {
			require.Equal(t, tc.want, isRobotBounded(tc.instructions))
		})
	}
}

func isRobotBounded(instructions string) bool {
	// A robot which is 'bounded' ends up in the same position after either 1, 2,
	// or 4 iterations of the instructions,

	type coord struct{ x, y int }

	pos := coord{0, 0}
	dir := 0
	move := func() {
		switch dir {
		case 0:
			pos.x++
		case 1:
			pos.y++
		case 2:
			pos.x--
		case 3:
			pos.y--
		}
	}
	start := coord{0, 0}
	for i := 0; i < 4; i++ {
		for _, inst := range instructions {
			switch inst {
			case 'G':
				move()
			case 'L':
				dir = (dir + 4 - 1) % 4
			case 'R':
				dir = (dir + 1) % 4
			}
		}
		if pos == start && dir == 0 {
			return true
		}
	}
	return false
}

package p0841keysandrooms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canVisitAllRooms(t *testing.T) {
	for _, tc := range []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rooms), func(t *testing.T) {
			require.Equal(t, tc.want, canVisitAllRooms(tc.rooms))
		})
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	seen := make([]bool, n)
	seen[0] = true
	todo := rooms[0]
	for len(todo) > 0 {
		roomKey := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		if seen[roomKey] {
			continue
		}
		seen[roomKey] = true
		todo = append(todo, rooms[roomKey]...)
	}

	for _, v := range seen {
		if !v {
			return false
		}
	}

	return true
}

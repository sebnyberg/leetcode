package p1847closestroom

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closestRoom(t *testing.T) {
	for _, tc := range []struct {
		rooms   [][]int
		queries [][]int
		want    []int
	}{
		{[][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}}, []int{3, -1, 3}},
		{[][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}}, []int{2, 1, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rooms), func(t *testing.T) {
			require.Equal(t, tc.want, closestRoom(tc.rooms, tc.queries))
		})
	}
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	// rooms[i] = [roomId, size]
	// queries[i] = [preferred, minSize]
	//
	// Sort queries by size
	// Sort rooms by size
	// For each query, add all rooms greater than or equal to the queried room size
	// Search for smallest abs distance

	// Sort rooms by size
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] > rooms[j][1]
	})

	k := len(queries)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	// Sort queries by size
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] > queries[j][1]
	})

	res := make([]int, k)
	var i int
	bigEnoughRooms := make([]int, 0, len(rooms))
	for _, q := range queries {
		prefer, minSize, idx := q[0], q[1], q[2]

		// add rooms until there are no more rooms big enough for this query
		for i < len(rooms) && rooms[i][1] >= minSize {
			bigEnoughRooms = append(bigEnoughRooms, rooms[i][0])
			i++
		}
		minDiff := math.MaxInt32
		minID := math.MaxInt32
		for _, roomID := range bigEnoughRooms {
			d := abs(roomID - prefer)
			if d == minDiff {
				if roomID < minID {
					minID = roomID
					minDiff = d
				}
			} else if d < minDiff {
				minID = roomID
				minDiff = d
			}
		}
		if minID == math.MaxInt32 {
			res[idx] = -1
		} else {
			res[idx] = minID
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package p2037minimumnumberofmovestoseateveryone

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMovesToSeat(t *testing.T) {
	for _, tc := range []struct {
		seats    []int
		students []int
		want     int
	}{
		{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
		{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
		{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.seats), func(t *testing.T) {
			require.Equal(t, tc.want, minMovesToSeat(tc.seats, tc.students))
		})
	}
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var diff int
	for i := range seats {
		diff += abs(seats[i] - students[i])
	}
	return diff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

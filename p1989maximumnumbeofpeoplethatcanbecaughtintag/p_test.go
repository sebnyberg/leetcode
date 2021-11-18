package p1989maximumnumbeofpeoplethatcanbecaughtintafunc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_catchMaximumAmountofPeople(t *testing.T) {
	for _, tc := range []struct {
		team []int
		dist int
		want int
	}{
		{[]int{0, 1, 0, 1, 0}, 3, 2},
		{[]int{1}, 1, 0},
		{[]int{0}, 1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.team), func(t *testing.T) {
			require.Equal(t, tc.want, catchMaximumAmountofPeople(tc.team, tc.dist))
		})
	}
}

func catchMaximumAmountofPeople(team []int, dist int) int {
	// The greedy hypothesis is that maximizing the distance between the person
	// who is "it" and a person who is not "it" will yield the best result.

	// Read from left-to-right - when encountering someone who is it, make this
	// person available

	// This feels like a simple window /stack exercise.
	// Read using a "fast" pointer until reaching the dist-1 position.
	//
}

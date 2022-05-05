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
	// Greedy solution
	// Add tagged and untagged players to one queue each.
	// If there is an untagged and tagged player within distance of the current
	// position, pop the most distant ones.
	untagged := make([]int, 0, len(team))
	tagged := make([]int, 0, len(team))

	var res int
	for i, player := range team {
		for len(untagged) > 0 && i-untagged[0] > dist {
			untagged = untagged[1:]
		}
		for len(tagged) > 0 && i-tagged[0] > dist {
			tagged = tagged[1:]
		}
		if player == 1 {
			tagged = append(tagged, i)
		} else {
			untagged = append(untagged, i)
		}
		if len(tagged) > 0 && len(untagged) > 0 {
			untagged = untagged[1:]
			tagged = tagged[1:]
			res++
		}
	}

	return res
}

package p1927sumgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumGame(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want bool
	}{
		{"?0?91172275656701?361205452?62??99?9??4478?7967373994600735??4?079246???5827572?81087461?089", true},
		{"25??", true},
		{"5023", false},
		{"?3295???", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, sumGame(tc.num))
		})
	}
}

func sumGame(num string) bool {
	n := len(num)
	var sums [2]int
	var moves [2]int
	for i, ch := range num {
		if ch == '?' {
			moves[i/(n/2)]++
		} else {
			sums[i/(n/2)] += int(ch - '0')
		}
	}

	// Alice picks the last number and wins
	if (moves[0]+moves[1])%2 == 1 {
		return true
	}

	// For the side that has more `?` than the other,
	// Alice wins if:
	// 1. That side's sum is greater than the other.
	// 2. The distance between sums is not 9*(# moves for Bob)
	movesForBob := abs(moves[0]-moves[1]) / 2
	sumDiff := abs(sums[1] - sums[0])
	if moves[0] > moves[1] {
		return sums[0] > sums[1] || sumDiff != 9*movesForBob
	}
	return sums[1] > sums[0] || sumDiff != 9*movesForBob
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// func sumGame(num string) bool {
// 	n := len(num)
// 	var sums [2]int
// 	var markCount [2]int
// 	for i, ch := range num {
// 		if ch == '?' {
// 			markCount[i/(n/2)]++
// 		} else {
// 			sums[i/(n/2)] += int(ch - '0')
// 		}
// 	}

// 	// If Alice has the final move, she will always win
// 	if (markCount[0]+markCount[1])%2 == 1 {
// 		return true
// 	}
// 	// If number of '?'s on both sides are equal, Bob wins
// 	if markCount[0] == markCount[1] {
// 		return sums[0] != sums[1]
// 	}

// 	// For even marks, figuring out who wins is a question of whether Bob can
// 	// bridge / keep the gap, or if Alice can extend the gap between the sums.
// 	// For any marks that are on both sides, the status remains unchanged.
// 	// For example, if Alice increases a high sum by adding '9', then Bob counters
// 	// with a '0' on the other side, and vice versa.
// 	// What remains is questionmarks on a single side, and the two sums.
// 	sum, marks, other := sums[0], markCount[0]-markCount[1], sums[1]
// 	if markCount[1] > markCount[0] {
// 		sum, marks, other = sums[1], markCount[1]-markCount[0], sums[0]
// 	}

// 	// At this point, Alice has the same amount of questionmarks as Bob.
// 	// Alice wants to either: a) keep the current sum below the other sum by
// 	// placing small numbers, or b) increase the current sum above the other sum.
// 	// She has marks/2 9's at her disposal.
// 	if sum < other {
// 		return sum+(marks/2)*9 > other || // place '9's
// 			other-sum > (marks/2)*9 // bob can't place enough 9's (Alice puts zeroes)
// 	}
// 	return true
// }

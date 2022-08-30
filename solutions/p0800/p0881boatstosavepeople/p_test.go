package p881boatstosavepeople

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numRescueBoats(t *testing.T) {
	for _, tc := range []struct {
		people []int
		limit  int
		want   int
	}{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.people), func(t *testing.T) {
			require.Equal(t, tc.want, numRescueBoats(tc.people, tc.limit))
		})
	}
}

func numRescueBoats(people []int, limit int) int {
	// Given constraints, we are looking for an O(n) or O(nlogn) solution

	// The first question to ask, is whether knowing how many boats we have can
	// simplify the algorithm. I.e., can we guess x boats, then verify or refute
	// that claim in O(n)?

	// First, for O(n), is there a greedy algorithm that will work?

	// Since there are only two people per boat, each person must belong to a
	// boat.

	// Consider matching the lowest weight and highest weight person. If the pair
	// has a limit greater than the total, then there was no matching that
	// could've worked; the lowest weight person was already too high weight for
	// a matching to succeed. Therefore, the highest weight person must sit in
	// their own boat.

	// This matching procedure can be executed until there are no more valid
	// matchings.

	// This solution is worst case O(nlogn) due to sorting (but practically O(n)
	// amortized for quicksort).

	sort.Ints(people)
	l, r := 0, len(people)-1
	var count int
	for l < r {
		if people[l]+people[r] > limit {
			r--
		} else {
			r--
			l++
		}
		count++
	}
	if l == r {
		count++
	}
	return count
}

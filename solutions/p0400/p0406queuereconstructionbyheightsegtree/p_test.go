package p0406queuereconstructionbyheightsegtree

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reconstructQueue(t *testing.T) {
	for _, tc := range []struct {
		people [][]int
		want   [][]int
	}{
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}, [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.people), func(t *testing.T) {
			require.Equal(t, tc.want, reconstructQueue(tc.people))
		})
	}
}

func reconstructQueue(people [][]int) [][]int {
	// Sort people by height ascending, number of people descending
	sort.Slice(people, func(i, j int) bool {
		a := people[i]
		b := people[j]
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})

	// Place people in order. The number of empty slots in front of each person
	// must match the number of taller or equal height people that is in front
	// of the person.
	//
	// For example, given (4,4), any next person in the list will be of height 4
	// or larger. So if we leave 4 empty slots in front of (4, 4) then we must
	// satisfy the requirement.
	//
	// The crux is to realise that adding people from high to low people in
	// front guarantees that the invariant holds.
	//
	// For this solution I will use a segtree to count number of empty slots in
	// the result.
	m := len(people)
	n := 1
	for n < m {
		n *= 2
	}

	// segtree is a segment-tree which keeps track of the number of slots for
	// different ranges in the tree.
	segtree := make([]int, n*2)
	for i := range people {
		segtree[n+i] = 1 // empty slot
	}
	for i := n - 1; i >= 1; i-- {
		segtree[i] = segtree[2*i] + segtree[2*i+1]
	}
	res := make([][]int, m)
	for i := range people {
		// find position of n empty slots
		want := people[i][1] + 1 // one extra slot to fit current person
		j := 1
		for j < n {
			l := segtree[j*2]
			if l >= want {
				j = j * 2
			} else {
				want -= l
				j = j*2 + 1
			}
		}

		// mark slot as full
		res[j-n] = people[i]

		// update segtree
		for k := j; k >= 1; k /= 2 {
			segtree[k]--
		}
	}
	return res
}

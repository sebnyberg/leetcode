package p0406queuereconstructionbyheight

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
	n := len(people)
	// Sort by height, position ascending
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	// Process people from small to tall.
	// Since all 'unprocessed' people are taller than the current, requirements in
	// terms of 'k' can be met by leaving empty slots in the result slice.
	res := make([][]int, n)
	for _, p := range people {
		var insertPos int
		var skipped int
		for len(res[insertPos]) != 0 || skipped < p[1] {
			if len(res[insertPos]) == 0 {
				skipped++
			}
			insertPos++
		}
		res[insertPos] = p
	}
	return res
}

package p1899mergetripletstoformtargettriplet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeTriplets(t *testing.T) {
	for _, tc := range []struct {
		triplets [][]int
		target   []int
		want     bool
	}{
		{[][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}, true},
		{[][]int{{1, 3, 4}, {2, 5, 8}}, []int{2, 5, 8}, true},
		{[][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}, true},
		{[][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.triplets), func(t *testing.T) {
			require.Equal(t, tc.want, mergeTriplets(tc.triplets, tc.target))
		})
	}
}

func mergeTriplets(triplets [][]int, target []int) bool {
	// Any merged triplet must be <= all elements in the target
	res := []int{0, 0, 0}
	for _, triplet := range triplets {
		for i := 0; i < 3; i++ {
			if triplet[i] > target[i] {
				goto ContinueSearch
			}
		}
		// Merge res with triplet
		for i := 0; i < 3; i++ {
			if res[i] < triplet[i] {
				res[i] = triplet[i]
			}
		}

	ContinueSearch:
	}
	for i := range res {
		if res[i] != target[i] {
			return false
		}
	}
	return true
}

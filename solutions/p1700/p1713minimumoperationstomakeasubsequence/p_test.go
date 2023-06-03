package p1713minimumoperationstomakeasubsequence

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		target []int
		arr    []int
		want   int
	}{
		{
			[]int{6, 4, 8, 1, 3, 2},
			[]int{4, 7, 6, 2, 3, 8, 6, 1},
			3,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.target, tc.arr))
		})
	}
}

func minOperations(target []int, arr []int) int {
	// Since numbers in target are unique, if arr[i] == target[j] then we know
	// that the matching is in position j in target (unless skipping the number,
	// of course).
	//
	// Consider a baseline cost of matching numbers if nothing exists in arr:
	//
	// cost = [ 1, 2, 3, 4 ]
	//
	// What happens to the cost when a number is matched?
	//
	// target = [ 1, 2, 3, 4 ]
	// arr    = [ 6, 3, 7, 2, 9 ]
	//               ^     ^
	//
	// When the 2 is encountered in arr, we could add the other numbers, giving
	// a cost of:
	//
	// cost = [ 1, 2, 2, 3 ]
	//
	// When the 2 is encountered, then we have a similar scenario, only we
	// are given the possibility of a better result: a total cost reduction of
	// 2.
	//
	// cost = [ 1, 1, 2, 3 ]
	//
	// It is clear that the 2 has more potential to get a better end result, but
	// could possibly reduce the total length of a matching sequence.
	//
	// For example, the above example clearly shows how matching the "1" in the
	// final position has a lot of "potential", but is not the optimal solution:
	//
	// target = [ 1, 2, 3, 4 ]
	// nums = [ 2, 3, 4, 1 ]
	//
	// This tells us that we are looking for the longest increasing subsequence
	// of matched indices in target. This is a famous problem (LIS) and can be
	// solved in O(n*logn) using various methods. I personally prefer Patience
	// sort.
	//
	numIdx := make(map[int]int)
	for i, x := range target {
		numIdx[x] = i
	}

	// Note: this step is not necessary. We could patience sort while reading
	// indices. I'm doing this for ease of understanding for someone visiting
	// this code.
	indices := []int{}
	for _, x := range arr {
		if i, exists := numIdx[x]; exists {
			indices = append(indices, i)
		}
	}
	costReduction := longestIncreasingSubsequence(indices)
	return len(target) - costReduction
}

func longestIncreasingSubsequence(indices []int) int {
	piles := []int{}
	for _, i := range indices {
		j := sort.SearchInts(piles, i)
		if j == len(piles) {
			piles = append(piles, j)
		}
		if piles[j] == i {
			continue
		}
		piles[j] = i
	}
	return len(piles)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

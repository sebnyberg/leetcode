package p0870advantageshuffle

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_advantageCount(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		B    []int
		want []int
	}{
		{[]int{12, 24, 8, 32}, []int{13, 25, 32, 11}, []int{24, 32, 8, 12}},
		{[]int{2, 7, 11, 15}, []int{1, 10, 4, 11}, []int{2, 11, 7, 15}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.A, tc.B), func(t *testing.T) {
			require.Equal(t, tc.want, advantageCount(tc.A, tc.B))
		})
	}
}

func advantageCount(A []int, B []int) []int {
	// A high number in A is a "high value number"
	// So, if we sort numbers in A by size (desc), the first occurrence will
	// be the most valuable number in A.
	// Similarly, a high value in B is difficult to beat, and low value
	// numbers in B are easy to beat.

	// Could use insertion sort here instead
	bb := make([]int, len(B))
	copy(bb, B)

	sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })
	sort.Slice(bb, func(i, j int) bool { return bb[i] > bb[j] })

	// Count values in A
	counts := make(map[int]int)
	for _, n := range A {
		counts[n]++
	}

	// Match values in the sorted versions of A and B
	matchedValues := make(map[int][]int)
	var j int
	for _, n := range bb {
		switch {
		case A[j] > n: // match - shift both i and j
			matchedValues[n] = append(matchedValues[n], A[j])
			counts[A[j]]--
			j++
		case A[j] <= n: // no match - do not shift j
		}
	}

	unmatchedStack := make([]int, 0)
	for unmatchedValue, count := range counts {
		for i := 0; i < count; i++ {
			unmatchedStack = append(unmatchedStack, unmatchedValue)
		}
	}

	// For values that matched, put them in the resulting array
	for i, n := range B {
		if matched := matchedValues[n]; len(matched) > 0 {
			A[i] = matched[len(matched)-1]
			matchedValues[n] = matched[:len(matched)-1] // Pop
			continue
		}
		// did not match value, pick some element from unmatched values
		A[i] = unmatchedStack[len(unmatchedStack)-1]
		unmatchedStack = unmatchedStack[:len(unmatchedStack)-1]
	}

	return A
}

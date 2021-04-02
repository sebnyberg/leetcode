package p0975oddevenjump

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_oddEvenJumps(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{10, 13, 12, 14, 15}, 2},
		{[]int{2, 3, 1, 1, 4}, 3},
		{[]int{27, 81, 84, 89, 58, 94, 57, 45, 66, 99}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, oddEvenJumps(tc.arr))
		})
	}
}

func oddEvenJumps(arr []int) int {
	// Worst description.... had to re-do this multiple times due to misunderstanding
	//
	// You start at a starting position i < n-1
	// You may only jump forward
	//
	// In the first jump, find the smallest possible number nums[j] >= nums[i]
	// such that j > i. If there are multiple candidates, pick the lowest one.
	//
	// In the second jump, the same rules apply, but instead find the greatest
	// possible number nums[j] <= nums[i].
	//
	// For jumps 3, 5, ..., the same rules apply as in the first jump.
	// Similarly for jump 4, 6, ..., the same rules apply as in the second jump.
	//
	//
	leadsToFinish := make(map[jumpPosition]bool)
	validStartPositions := 0
	n := len(arr)
	minFinder := NewIdxFinder(arr, func(a, b int) int {
		if a < b {
			return -1
		} else if a == b {
			return 0
		}
		return 1
	})
	maxFinder := NewIdxFinder(arr, func(a, b int) int {
		if a > b {
			return -1
		} else if a == b {
			return 0
		}
		return 1
	})

	for i := range arr {
		start := jumpPosition{i, false}
		if explorePosition(start, leadsToFinish, n, minFinder, maxFinder) {
			validStartPositions++
		}
	}

	return validStartPositions
}

type jumpPosition struct {
	i    int
	even bool
}

func explorePosition(pos jumpPosition, leadsToFinish map[jumpPosition]bool, n int, minFinder *IdxFinder, maxFinder *IdxFinder) bool {
	if pos.i == n-1 {
		return true
	}
	if finish, exists := leadsToFinish[pos]; exists && finish {
		return finish
	}
	var jumpTo int
	if !pos.even {
		// Smallest number >= nums[pos.i]
		jumpTo = minFinder.FindMinIdx(pos.i)
	} else {
		// Greatest number <= nums[pos.i]
		jumpTo = maxFinder.FindMinIdx(pos.i)
	}
	if jumpTo == -1 {
		leadsToFinish[pos] = false
		return false
	}
	nextPos := jumpPosition{jumpTo, !pos.even}
	leadsToFinish[pos] = explorePosition(nextPos, leadsToFinish, n, minFinder, maxFinder)
	return leadsToFinish[pos]
}

type SortedVal struct {
	idx int
	val int
}

type IdxFinder struct {
	origToSortedIdx []int
	sortedVals      []SortedVal
}

func NewIdxFinder(arr []int, comp func(a, b int) int) *IdxFinder {
	var f IdxFinder
	sortedVals := make([]SortedVal, len(arr))
	for i, n := range arr {
		sortedVals[i] = SortedVal{i, n}
	}
	sort.Slice(sortedVals, func(i, j int) bool {
		if comp(sortedVals[i].val, sortedVals[j].val) == 0 {
			return sortedVals[i].idx < sortedVals[j].idx
		}
		return comp(sortedVals[i].val, sortedVals[j].val) == -1
	})
	f.sortedVals = sortedVals
	f.origToSortedIdx = make([]int, len(sortedVals))
	for i, v := range sortedVals {
		f.origToSortedIdx[v.idx] = i
	}
	return &f
}

func (f *IdxFinder) FindMinIdx(i int) int {
	sortIdx := f.origToSortedIdx[i]
	curIdx := sortIdx + 1
	for curIdx < len(f.sortedVals) {
		if f.sortedVals[curIdx].idx > i {
			return f.sortedVals[curIdx].idx
		}
		curIdx++
	}
	return -1
}

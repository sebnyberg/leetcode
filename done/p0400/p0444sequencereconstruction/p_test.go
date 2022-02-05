package p0444sequencereconstruction

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sequenceReconstruction(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		sequences [][]int
		want      bool
	}{
		{[]int{1, 2, 3}, leetcode.ParseMatrix("[[1,2],[1,3]]"), false},
		{[]int{1, 2, 3}, leetcode.ParseMatrix("[[1,2]]"), false},
		{[]int{1, 2, 3}, leetcode.ParseMatrix("[[1,2],[1,3],[2,3]]"), true},
	} {
		t.Run(fmt.Sprintf("%+v,%+v", tc.nums, tc.sequences), func(t *testing.T) {
			require.Equal(t, tc.want, sequenceReconstruction(tc.nums, tc.sequences))
		})
	}
}

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	// seqNum points to the next element in a sequence
	type seqNum struct {
		seqIdx int
		idx    int
	}

	// adj contains a list of next elements in sequences
	adj := make([][]seqNum, n+1)
	indeg := make([]int, n+1)
	for i, seq := range sequences {
		for j := 0; j < len(seq)-1; j++ {
			adj[seq[j]] = append(adj[seq[j]], seqNum{seqIdx: i, idx: j + 1})
			indeg[seq[j+1]]++
		}
	}

	// Find initial element with an indegree of zero
	cur := -1
	for i := 1; i < len(indeg); i++ {
		if indeg[i] == 0 {
			if cur != -1 {
				return false
			}
			cur = i
		}
	}

	var numPos int
	var next int
	for cur != -1 {
		next = -1

		// Visit adjacent numbers and reduce indegree. If there are two
		// alternatives, then there are more than one valid solutions.
		for _, num := range adj[cur] {
			x := sequences[num.seqIdx][num.idx]
			indeg[x]--
			if indeg[x] == 0 {
				if next != -1 {
					return false
				}
				next = x
			}
		}
		numPos++
		cur = next
	}

	return numPos == n
}

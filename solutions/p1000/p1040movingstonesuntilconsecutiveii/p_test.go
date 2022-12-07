package p1040movingstonesuntilconsecutiveii

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numMovesStonesII(t *testing.T) {
	for i, tc := range []struct {
		stones []int
		want   []int
	}{
		{[]int{7, 4, 9}, []int{1, 2}},
		{[]int{6, 5, 4, 3, 10}, []int{2, 3}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numMovesStonesII(tc.stones))
		})
	}
}

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	n := len(stones)

	// Sort out annoying edge-case
	if stones[n-1]-stones[0]+1 == n {
		return []int{0, 0}
	}

	// Find a range of stones that cover n in total
	var l int
	minMoves := math.MaxInt32
	for r := 1; r < len(stones); r++ {
		for stones[r]-stones[l]+1 > n {
			l++
		}
		// [l,r] now covers <= n
		// The number of stones in the interval is the number of "free" moves
		k := n - (r - l + 1)
		// Except! There is an annoying case where it seems as though we need to
		// move one stone, but due to the "no endpoint to enpoint" condition, we
		// actually need two moves...
		if r-l+1 == n-1 && ((l == 0 && stones[l] == stones[l+1]-1) || r == n-1 && stones[r] == stones[r-1]+1) {
			minMoves = min(minMoves, 2)
		} else {
			minMoves = min(minMoves, k)
		}
	}

	// Moving an endpoint reduces the range of possible moves by the number of
	// empty slots between the current position and the next position.
	// The best way to minimize the loss is to move the current stone to a
	// position that is just after the next stone.
	// No matter how we do this, we will lose the moves in the gap between the
	// first and second stone, or penultimate and ultimate stone.
	var maxMoves int
	if stones[1]-stones[0] < stones[n-1]-stones[n-2] {
		maxMoves = stones[n-1] - stones[1] - n + 2
	} else {
		maxMoves = stones[n-2] - stones[0] - n + 2
	}

	return []int{minMoves, maxMoves}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

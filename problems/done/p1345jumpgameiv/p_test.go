package p1345jumpgameiv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minJumps(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3},
		{[]int{7}, 0},
		{[]int{7, 6, 9, 6, 9, 6, 9, 7}, 1},
		{[]int{6, 1, 9}, 2},
		{[]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, minJumps(tc.in))
		})
	}
}

func Test_minJumpsBig(t *testing.T) {
	in := make([]int, 1e5)
	for i := 0; i < 1e5; i++ {
		in[i] = 7
	}
	require.Equal(t, 1, minJumps(in))
}

func minJumps(arr []int) int {
	if len(arr) <= 2 {
		return len(arr) - 1
	}

	// Collect jump indices for each integer in the list
	jumps := make(map[int][]int)
	for i, n := range arr {
		jumps[n] = append(jumps[n], i)
	}

	// Traverse the list, branching out when possible
	// and marknig visited nodes in th eseen map
	cur := []int{0}
	var next []int
	seen := make([]bool, len(arr))
	moves := 0

	for {
		// Reset
		next = make([]int, 0)

		// Explore options at the current indices
		for _, idx := range cur {
			if idx == len(arr)-1 { // Done
				return moves
			}

			// Move a step forward
			if !seen[idx+1] {
				next = append(next, idx+1)
				seen[idx+1] = true
			}

			// Move to all jump indices
			for _, jumpidx := range jumps[arr[idx]] {
				if !seen[jumpidx] {
					next = append(next, jumpidx)
					seen[jumpidx] = true
				}
			}
			delete(jumps, arr[idx])

			// Move a step backward
			if idx > 0 && !seen[idx-1] {
				seen[idx-1] = true
				next = append(next, idx-1)
			}
		}

		moves++
		cur = next
	}
}

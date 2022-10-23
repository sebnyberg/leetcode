package p2440createcomponentswithsamevalue

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_componentValue(t *testing.T) {
	for i, tc := range []struct {
		nums  []int
		edges [][]int
		want  int
	}{
		{[]int{1, 1, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 2}, leetcode.ParseMatrix("[[12,14],[14,8],[8,4],[4,16],[16,1],[1,15],[15,5],[5,6],[6,9],[9,10],[10,17],[17,19],[19,13],[13,0],[0,2],[2,3],[3,18],[18,7],[7,11]]"), 1},
		{[]int{2}, [][]int{}, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, componentValue(tc.nums, tc.edges))
		})
	}
}

func componentValue(nums []int, edges [][]int) int {
	var sum int
	numsCpy := make([]int, len(nums))
	for _, x := range nums {
		sum += x
	}
	adj := make([][]int, len(nums))
	deg := make([]int, len(nums))
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
		deg[e[0]]++
		deg[e[1]]++
	}
	degCpy := make([]int, len(nums))
	curr := []int{}
	next := []int{}
	reset := func() {
		curr = curr[:0]
		next = next[:0]
		for i := range nums {
			numsCpy[i] = nums[i]
			degCpy[i] = deg[i]
			if deg[i] == 1 {
				curr = append(curr, i)
			}
		}
	}
	try := func(nums, deg []int, adj [][]int, k int) int {
		var cuts int
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				if nums[x] > k {
					return -1
				}
			}
			for _, x := range curr {
				if nums[x] > k {
					return -1
				}
				if deg[x] == 0 {
					if nums[x] == k {
						return cuts
					}
				}
				for _, nei := range adj[x] {
					if deg[nei] == 0 {
						continue
					}
					if nums[x] == k {
						// cut
						cuts++
					} else {
						// must merge current with next
						nums[nei] += nums[x]
					}
					deg[x]--
					deg[nei]--
					if deg[nei] == 1 {
						next = append(next, nei)
					}
				}
			}
			curr, next = next, curr
		}
		return cuts
	}

	for k := 1; k <= sum; k++ {
		if sum%k != 0 {
			continue
		}
		reset()
		if d := try(numsCpy, degCpy, adj, k); d >= 0 {
			return d
		}
	}
	return 0
}

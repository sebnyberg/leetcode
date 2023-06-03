package p1707maximuxorwithanelementfromarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maximizeXor(t *testing.T) {
	for i, tc := range []struct {
		nums    []int
		queries [][]int
		want    []int
	}{
		{[]int{5, 2, 4, 6, 6, 3}, leetcode.ParseMatrix("[[12,4],[8,1],[6,3]]"), []int{15, -1, 5}},
		{[]int{0, 1, 2, 3, 4}, leetcode.ParseMatrix("[[3,1],[1,3],[5,6]]"), []int{3, 3, 7}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximizeXor(tc.nums, tc.queries))
		})
	}
}

type trieNode struct {
	next [2]*trieNode
}

func maximizeXor(nums []int, queries [][]int) []int {
	// The solution to this problem is quite intuitive.
	//
	// 1. Sort queries by max value
	// 2. Create a trie of numbers where each node denotes the presence of a
	// bit from most significant (MSB) to least significant (LSB).
	// 3. For each query, put all numbers <= max into trie. Then search for the
	// longest prefix that maximizes the XOR sum. NOTE! The reason why this is
	// optimal is because any matching XOR MSB is by definition worth more than
	// one unmathed MSB and the rest matching. For examble, matching 1111 to
	// 0111 (1111 XOR 0111 = 1000) is better than matching 1111 to 1000
	// (1111 XOR 1000 = 0111).
	m := len(queries)
	sort.Ints(nums)

	// Sort queries by maximum number
	// idx[i] = original index of query in sorted position i
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return queries[idx[i]][1] < queries[idx[j]][1]
	})

	root := &trieNode{}

	var j int
	results := make([]int, m)
	for i := range idx {
		q := queries[idx[i]]
		for ; j < len(nums) && nums[j] <= q[1]; j++ {
			// Add number to trie
			x := nums[j]
			curr := root
			for msb := 30; msb >= 0; msb-- {
				next := (x >> msb) & 1
				if curr.next[next] == nil {
					curr.next[next] = &trieNode{}
				}
				curr = curr.next[next]
			}
		}
		if j == 0 {
			results[idx[i]] = -1
			continue
		}
		// Search for maximum common prefix of inverse of the number in the
		// query.
		x := ^uint(q[0])
		curr := root
		var res int
		for msb := 30; msb >= 0 && curr != nil; msb-- {
			want := (x >> msb) & 1
			if curr.next[want] != nil {
				res += (1 << msb)
				curr = curr.next[want]
			} else {
				curr = curr.next[1-want]
			}
		}
		results[idx[i]] = res
	}
	return results
}

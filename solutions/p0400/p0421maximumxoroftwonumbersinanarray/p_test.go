package p0421maximumxoroftwonumbersinanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaximumXOR(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 10, 5, 25, 2, 8}, 28},
		{[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, 127},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaximumXOR(tc.nums))
		})
	}
}

type trieNode struct {
	next [2]*trieNode
}

func findMaximumXOR(nums []int) int {
	// Any number which has a bit in a position so high that another number does
	// not have it must be included in the solution.
	// We can put all numbers into a trie then search through the trie for a path
	// with a maximum amount of ones given the current number
	root := &trieNode{}
	addToTrie := func(num int) {
		cur := root
		for i := 30; i >= 0; i-- {
			bm := 1 << i
			bit := num & bm >> i
			if cur.next[bit] == nil {
				cur.next[bit] = &trieNode{}
			}
			cur = cur.next[bit]
		}
	}
	addToTrie(nums[0])
	findMaxNum := func(num int) int {
		res := 0
		cur := root
		mask := 0xffffffff >> 1
		want := ^num & mask
		for i := 30; i >= 0; i-- {
			bm := 1 << i
			wantBit := (want & bm) >> i
			if cur.next[wantBit] != nil {
				cur = cur.next[wantBit]
				res |= bm
			} else {
				cur = cur.next[^wantBit&1]
			}
		}
		return res
	}

	var maxRes int
	for _, num := range nums[1:] {
		// Search trie to find maximum number combined with this one
		maxRes = max(maxRes, findMaxNum(num))
		addToTrie(num)
	}

	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

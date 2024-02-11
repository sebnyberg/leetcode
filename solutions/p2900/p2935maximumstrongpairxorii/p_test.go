package p2935maximumstrongpairxorii

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumStrongPairXor(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumStrongPairXor(tc.nums))
		})
	}
}

type trieNode struct {
	next  [2]*trieNode
	count int
}

const maxBits = 21

func maximumStrongPairXor(nums []int) int {
	// There are a couple of ways to attack this problem.
	//
	// First, let's consider the relationship given by the equation.
	//
	// We can think of it as the current smallest and current largest value.
	// The possible ranges to consider are:
	//
	// [1,2], [2,4], [3,6], [4,8], ...
	//
	// That is,
	// If the first number is 1, then any other number in the range [1,2] is ok.
	// If the first number is 2, then any other number in the range [2,4] is ok.
	//
	// If we can, given a first number, efficiently compare that number to all
	// available numbers in the acceptable range, then we can quickly find the
	// answer to the question.
	//
	// Usually for XOR problems, the answer involves a Trie somehow. We can
	// clearly efficiently keep a range of numbers within the trie, but I wonder
	// how having the numbers in the trie can help us search for the highest XOR
	// pair... Oh.. I see. The thing is that combining numbers to get high bits
	// is always preferable no matter what other bits are set. That is, we
	// should always prefer mismatching bits whenever possible in the trie.
	//
	// Alright, that's all I need to know!

	root := &trieNode{}
	addToTrie := func(x int) {
		curr := root
		root.count++
		for i := maxBits; i >= 0; i-- {
			b := (x >> i) & 1
			if curr.next[b] == nil {
				curr.next[b] = &trieNode{}
			}
			curr = curr.next[b]
			curr.count++
		}
	}

	removeFromTrie := func(x int) {
		curr := root
		for i := maxBits; i >= 0; i-- {
			b := (x >> i) & 1
			curr = curr.next[b]
			curr.count--
		}
	}

	// First, clear out duplicates
	sort.Ints(nums)
	var j int
	for i := range nums {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	nums = nums[:j+1]

	// The trie functions seem to be working. Now the idea is to, for each
	// unique number, add all numbers above until no more numbers can be added.
	// Then remove that unique number from the trie and continue.
	var res int
	j = 0
	for i := range nums {
		for j < len(nums) && nums[j]-nums[i] <= nums[i] {
			addToTrie(nums[j])
			j++
		}
		// Find best matched pair
		curr := root
		var result int
		if root.count > 0 {
			for k := maxBits; k >= 0; k-- {
				b := (nums[i] >> k) & 1
				if curr.next[1-b] != nil && curr.next[1-b].count > 0 {
					curr = curr.next[1-b]
					result |= (1 << k)
				} else {
					if curr.next[b] == nil || curr.next[b].count == 0 {
						panic("no options in trie, must be a bug")
					}
					curr = curr.next[b]
				}
			}
		}
		res = max(res, result)

		removeFromTrie(nums[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

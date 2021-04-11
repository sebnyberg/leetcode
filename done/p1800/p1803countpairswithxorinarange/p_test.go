package p1803countpairswithxorinarange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		low  int
		high int
		want int
	}{
		{[]int{1, 4, 2, 7}, 2, 6, 6},
		{[]int{9, 8, 4, 2, 1}, 5, 14, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countPairs(tc.nums, tc.low, tc.high))
		})
	}
}

func countPairs(nums []int, low int, high int) int {
	root := &TrieNode{}
	res := 0
	for _, n := range nums {
		res += root.countLess(n, high+1, 14) - root.countLess(n, low, 14)
		root.insert(n, 14)
	}
	return res
}

type TrieNode struct {
	next  [2]*TrieNode
	count int
}

func (n *TrieNode) insert(num, i int) {
	n.count++
	if i < 0 {
		return
	}
	b := (num >> i) & 1
	if n.next[b] == nil {
		n.next[b] = &TrieNode{}
	}
	n.next[b].insert(num, i-1)
}

func (n *TrieNode) countLess(num, lim, i int) int {
	if i < 0 {
		return 0
	}
	b := (num >> i) & 1
	limB := (lim >> i) & 1
	res := 0
	// When the limit byte is 1
	// If there is a value XOR b that yields 0,
	// the count of that tree is added to res
	// Then the search continues along the path of equality (b == limit)
	switch {
	case b == 0 && limB == 0:
		if n.next[0] != nil {
			res += n.next[0].countLess(num, lim, i-1)
		}
	case b == 0 && limB == 1:
		if n.next[0] != nil {
			res += n.next[0].count
		}
		if n.next[1] != nil {
			res += n.next[1].countLess(num, lim, i-1)
		}
	case b == 1 && limB == 0:
		if n.next[1] != nil {
			res += n.next[1].countLess(num, lim, i-1)
		}
	case b == 1 && limB == 1:
		if n.next[1] != nil {
			res += n.next[1].count
		}
		if n.next[0] != nil {
			res += n.next[0].countLess(num, lim, i-1)
		}
	}
	return res
}

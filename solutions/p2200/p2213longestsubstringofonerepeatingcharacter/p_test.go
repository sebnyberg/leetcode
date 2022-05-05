package p2213longestsubstringofonerepeatingcharacter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestRepeating(t *testing.T) {
	for _, tc := range []struct {
		s               string
		queryCharacters string
		queryIndices    []int
		want            []int
	}{
		{"babacc", "bcb", []int{1, 3, 3}, []int{3, 3, 4}},
		{"abyzz", "aa", []int{2, 1}, []int{2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestRepeating(tc.s, tc.queryCharacters, tc.queryIndices))
		})
	}
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	k := len(queryIndices)
	tree := NewSegmentTree(s)
	res := make([]int, len(queryIndices))
	for i := 0; i < k; i++ {
		tree.update(0, 0, len(s)-1, queryIndices[i], queryCharacters[i])
		res[i] = tree.tree[0].maxVal
	}
	return res
}

type Node struct {
	maxVal                 int
	prefixStart, prefixEnd int
	suffixStart, suffixEnd int
}

type SegmentTree struct {
	tree []Node
	s    []byte
}

func NewSegmentTree(s string) *SegmentTree {
	t := &SegmentTree{
		tree: make([]Node, len(s)*4),
		s:    []byte(s),
	}
	t.build(0, 0, len(s)-1)
	return t
}

func (t *SegmentTree) build(pos, l, r int) {
	if l == r {
		t.tree[pos] = Node{1, l, l, r, r}
		return
	}
	mid := l + (r-l)/2

	t.build(2*pos+1, l, mid)
	t.build(2*pos+2, mid+1, r)

	t.tree[pos] = t.merge(t.tree[2*pos+1], t.tree[2*pos+2], l, mid, r)
}

func (t *SegmentTree) merge(left, right Node, l, mid, r int) Node {
	maxVal := max(left.maxVal, right.maxVal)

	prefixStart := left.prefixStart
	prefixEnd := left.prefixEnd
	suffixStart := right.suffixStart
	suffixEnd := right.suffixEnd

	if t.s[mid] == t.s[mid+1] {
		maxVal = max(maxVal, right.prefixEnd-left.suffixStart+1)

		if left.prefixEnd-left.prefixStart+1 == mid-l+1 {
			prefixEnd = right.prefixEnd
		}

		if right.suffixEnd-right.suffixStart+1 == r-mid {
			suffixStart = left.suffixStart
		}
	}
	return Node{
		maxVal:      maxVal,
		prefixStart: prefixStart,
		prefixEnd:   prefixEnd,
		suffixStart: suffixStart,
		suffixEnd:   suffixEnd,
	}
}

func (t *SegmentTree) update(pos, l, r, idx int, ch byte) {
	if l == r {
		t.tree[pos] = Node{1, l, l, r, r}
		t.s[idx] = ch
		return
	}

	mid := l + (r-l)/2
	if idx <= mid {
		t.update(2*pos+1, l, mid, idx, ch)
	} else {
		t.update(2*pos+2, mid+1, r, idx, ch)
	}
	t.tree[pos] = t.merge(t.tree[2*pos+1], t.tree[2*pos+2], l, mid, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

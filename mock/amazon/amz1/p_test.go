package amz1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTarget(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}

	res := findTarget(tree, 9)
	require.Equal(t, true, res)

	tree2 := &TreeNode{Val: 1}
	res = findTarget(tree2, 2)
	require.Equal(t, false, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The problem here is that the graph can contain negative nodes
// This means that we cannot cut off the right side of the BST even though
// k is greater than the current node value.. There could be a value which
// works anyway in the negative interval.
// Since the number of nodes is < 10000, I'll collect all numbers doing
// in-order traversal, then work with the list instead...
func findTarget(root *TreeNode, k int) bool {
	vals := make([]int, 0)
	collect(root, &vals)
	for i := range vals {
		for j := range vals[i+1:] {
			if vals[i]+vals[i+1+j] == k {
				return true
			}
		}
	}
	return false
}

func collect(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	collect(root.Left, vals)
	*vals = append(*vals, root.Val)
	collect(root.Right, vals)
}

func Test_letterCombinations(t *testing.T) {
	for _, tc := range []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.digits), func(t *testing.T) {
			require.Equal(t, tc.want, letterCombinations(tc.digits))
		})
	}
}

func letterCombinations(digits string) []string {
	var f combinationFinder
	f.findCombinations(digits, []rune{})
	return f.found
}

type combinationFinder struct {
	found []string
}

var letters = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func (f *combinationFinder) findCombinations(digits string, prefix []rune) {
	if len(digits) == 0 {
		if len(prefix) > 0 {
			f.found = append(f.found, string(prefix))
		}
		return
	}
	for i, r := range letters[rune(digits[0])] {
		var prefixCpy []rune
		if i == 0 {
			prefixCpy = prefix
		} else {
			prefixCpy = make([]rune, len(prefix))
			copy(prefixCpy, prefix)
		}
		prefixCpy = append(prefixCpy, r)
		f.findCombinations(digits[1:], prefixCpy)
	}
}

func Test_copyRandomList(t *testing.T) {
	nodes := make([]*Node, 5)
	for i, value := range []int{7, 13, 11, 10, 1} {
		nodes[i] = &Node{Val: value}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	for idx, target := range []*Node{
		nil, nodes[0], nodes[4], nodes[2], nodes[0],
	} {
		nodes[idx].Random = target
	}

	res := copyRandomList(nodes[0])
	require.NotNil(t, res)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Stage 1. Point next in original list to new nodes, and
	// random in new list to next
	cur := head
	for cur != nil {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}

	// Stage 2. Set random pointers
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// Stage 3. Fix next pointers
	copiedHead := head.Next
	cur = head
	for {
		if cur.Next.Next == nil {
			cur, cur.Next = nil, nil
			break
		}
		cur, cur.Next, cur.Next.Next = cur.Next.Next, cur.Next.Next, cur.Next.Next.Next
	}

	return copiedHead
}

func Test_trap(t *testing.T) {
	for _, tc := range []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.height), func(t *testing.T) {
			require.Equal(t, tc.want, trap(tc.height))
		})
	}
}

func trap(height []int) int {
	// 1. Find max point in heights
	var maxHeight, maxIdx int
	for i, h := range height {
		if h > maxHeight {
			maxIdx = i
			maxHeight = h
		}
	}

	// 2. Scan from both sides to maxIdx, adding diff between max so far and height
	var maxSoFar int
	var waterAmt int
	for i := 0; i < maxIdx; i++ {
		if height[i] >= maxSoFar {
			maxSoFar = height[i]
			continue
		}
		waterAmt += maxSoFar - height[i]
	}

	maxSoFar = 0
	for i := len(height) - 1; i > maxIdx; i-- {
		if height[i] >= maxSoFar {
			maxSoFar = height[i]
			continue
		}
		waterAmt += maxSoFar - height[i]
	}

	return waterAmt
}

package p0366findleavesofbinarytree

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTree(input string) *TreeNode {
	// Trim start/end []
	input = input[1 : len(input)-1]

	// Split by comma
	inputParts := strings.Split(input, ",")
	n := len(inputParts)

	if n == 0 || inputParts[0] == "" {
		return nil
	}

	// Create one node per element in the array
	nodes := make([]*TreeNode, n)
	for i, part := range inputParts {
		if part != "null" {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			nodes[i] = &TreeNode{Val: val}
		}
	}

	q := list.New()
	q.PushBack(nodes[0])

	i := 1
	for q.Len() > 0 && i < n {
		el := q.Remove(q.Front()).(*TreeNode)
		if nodes[i] != nil {
			el.Left = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
		if i >= n {
			break
		}
		if nodes[i] != nil {
			el.Right = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
	}

	return nodes[0]
}

func Test_findLeaves(t *testing.T) {
	for _, tc := range []struct {
		root string
		want [][]int
	}{
		{"[1,2,3,4,5]", [][]int{{4, 5, 3}, {2}, {1}}},
		{"[1]", [][]int{{1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.Equal(t, tc.want, findLeaves(root))
		})
	}
}

func findLeaves(root *TreeNode) [][]int {
	f := leavesFinder{
		leaves: make([][]int, 0),
	}
	f.subTreeHeight(root, 0)
	return f.leaves
}

type leavesFinder struct {
	leaves [][]int
}

// subTreeHeight returns the height of a subtree and adds the current node to
// its corresponding "leaves level" based on the max height of its subtrees.
func (f *leavesFinder) subTreeHeight(cur *TreeNode, curLevel uint8) int {
	if cur == nil {
		return 0
	}
	leftHeight := f.subTreeHeight(cur.Left, curLevel+1)
	rightHeight := f.subTreeHeight(cur.Right, curLevel+1)
	h := max(leftHeight, rightHeight)
	// Incease length of leaves slice as needed
	if d := h - len(f.leaves) + 1; d >= 1 {
		f.leaves = append(f.leaves, make([][]int, d)...)
	}
	f.leaves[h] = append(f.leaves[h], cur.Val)
	return h + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

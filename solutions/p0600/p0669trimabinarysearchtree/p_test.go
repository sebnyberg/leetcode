package p0669trimabinarysearchtree

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

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == nil || other == nil {
		return n == other
	}
	if n.Val != other.Val {
		return false
	}
	return n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
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

func Test_trimBST(t *testing.T) {
	for _, tc := range []struct {
		tree      string
		low, high int
		want      string
	}{
		{"[1,0,2]", 1, 2, "[1,null,2]"},
		{"[3,0,4,null,2,null,null,1]", 1, 3, "[3,2,null,1]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tt := ParseTree(tc.tree)
			want := ParseTree(tc.want)
			got := trimBST(tt, tc.low, tc.high)
			require.True(t, got.Equals(want))
		})
	}
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// First we 'need' to find the root of the trimmed tree, which may not be the
	// provided root.
	root = findRoot(root, low, high)
	if root == nil {
		return nil
	}

	// Then we may trim the tree based on some simple rules.

	// Visit each valid node.
	// For each valid node, if a left non-nil descendant is smaller than low,
	// replace that descendant with that descendant's right descendant.
	// If a right non-nil descendant is larger than high, replace that descendant
	// with its left descendant.
	// These operations are safe because of the search tree invariants. Draw this
	// on a piece of paper for a proof.

	trim(root, low, high)
	return root
}

func trim(cur *TreeNode, low, high int) {
	if cur == nil {
		return
	}
	if cur.Left != nil {
		if cur.Left.Val < low {
			if cur.Left.Right != nil {
				cur.Left = cur.Left.Right
			} else {
				cur.Left = nil
			}
			trim(cur, low, high)
			return
		}
	}
	if cur.Right != nil {
		if cur.Right.Val > high {
			if cur.Right.Left != nil {
				cur.Right = cur.Right.Left
			} else {
				cur.Right = nil
			}
			trim(cur, low, high)
		}
	}
	trim(cur.Left, low, high)
	trim(cur.Right, low, high)
}

func findRoot(cur *TreeNode, low, high int) *TreeNode {
	if cur == nil {
		return nil
	}
	if cur.Val >= low && cur.Val <= high {
		return cur
	}
	if cur.Val < low {
		return findRoot(cur.Right, low, high)
	}
	return findRoot(cur.Left, low, high)
}

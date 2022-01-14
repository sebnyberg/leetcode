package p0701insertintobinarysearchtree

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
	return n.Val == other.Val && n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
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

func Test_insertIntoBST(t *testing.T) {
	for _, tc := range []struct {
		root string
		val  int
		want string
	}{
		{"[4,2,7,1,3]", 5, "[4,2,7,1,3,5]"},
		{"[40,20,60,10,30,50,70]", 25, "[40,20,60,10,30,50,70,null,null,25]"},
		{"[4,2,7,1,3,null,null,null,null,null,null]", 5, "[4,2,7,1,3,5]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.True(t, insertIntoBST(root, tc.val).Equals(ParseTree(tc.want)))
		})
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// The invariant of the BST is that given a node's value, all subsequent
	// values are larger/smaller than the current one.
	if root == nil { // edge-case
		return &TreeNode{Val: val}
	}
	if root.Val > val { // Go left
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else { // root.Val < val
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}

package p0897increasingordersearchtree

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

func (n *TreeNode) FindVal(val int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val == val {
		return n
	}
	if l := n.Left.FindVal(val); l != nil {
		return l
	}
	return n.Right.FindVal(val)
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

func Test_increasingBST(t *testing.T) {
	for _, tc := range []struct {
		root string
		want string
	}{
		{"[5,3,6,2,4,null,8,1,null,null,null,7,9]", ""},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			res := increasingBST(tree)
			require.Equal(t, tc.want, res)
		})
	}
}

func increasingBST(root *TreeNode) *TreeNode {
	// We can perform reverse inorder and keep a reference to the next node
	// (to which the current node will be the new .Right)
	var next *TreeNode
	visit(root, &next)
	return next
}

func visit(cur *TreeNode, next **TreeNode) {
	if cur == nil {
		return
	}
	visit(cur.Right, next)
	cur.Right = *next
	*next = cur
	visit(cur.Left, next)
	cur.Left = nil
}

package p0572subtreeofanothertree

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_isSubtree(t *testing.T) {
	for _, tc := range []struct {
		root    string
		subRoot string
		want    bool
	}{
		{"[3,4,5,1,2]", "[4,1,2]", true},
		{"[3,4,5,1,2,null,null,null,null,0]", "[4,1,2]", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			subRoot := ParseTree(tc.subRoot)
			require.Equal(t, tc.want, isSubtree(root, subRoot))
		})
	}
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return visit(root, subRoot)
}

func visit(curr, subRoot *TreeNode) bool {
	if curr == nil {
		return false
	}

	if curr.Val == subRoot.Val {
		if match(curr, subRoot) {
			return true
		}
	}

	return visit(curr.Left, subRoot) || visit(curr.Right, subRoot)
}

func match(curr, sub *TreeNode) bool {
	if curr == nil && sub != nil || curr != nil && sub == nil {
		return false
	}
	if curr == nil && sub == nil {
		return true
	}
	if curr.Val != sub.Val {
		return false
	}
	return match(curr.Left, sub.Left) && match(curr.Right, sub.Right)
}

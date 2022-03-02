package p0662maximumwidthofbinarytree

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

func Test_widthOfBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[1,3,2,5]", 2},
		{"[1,3,2,5,3,null,9]", 4},
		{"[1,3,null,5,3]", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, widthOfBinaryTree(tree))
		})
	}
}

func widthOfBinaryTree(root *TreeNode) int {
	type nodeIdx struct {
		n   *TreeNode
		idx int
	}
	cur := []nodeIdx{{root, 1}}
	next := []nodeIdx{}
	var maxCount int
	for len(cur) > 0 {
		if d := cur[len(cur)-1].idx - cur[0].idx + 1; d > maxCount {
			maxCount = d
		}
		next = next[:0]
		var count int
		for _, n := range cur {
			count++
			if n.n.Left != nil {
				next = append(next, nodeIdx{n.n.Left, 2 * n.idx})
			}
			if n.n.Right != nil {
				next = append(next, nodeIdx{n.n.Right, 2*n.idx + 1})
			}
		}
		cur, next = next, cur
	}
	return maxCount
}

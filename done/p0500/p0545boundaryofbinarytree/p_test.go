package p0545boundaryofbinarytree

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

func Test_boundaryOfBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		root string
		want []int
	}{
		{"[1]", []int{1}},
		{"[1,2,3,4,5,6,null,null,null,7,8,9,10]", []int{1, 2, 4, 7, 8, 9, 10, 6, 3}},
		{"[1,null,2,3,4]", []int{1, 3, 4, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, boundaryOfBinaryTree(tree))
		})
	}
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	res := []int{root.Val}
	collectLeft(root.Left, &res)
	collectLeaves(root.Left, &res)
	collectLeaves(root.Right, &res)
	collectRight(root.Right, &res)
	return res
}

func collectLeft(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	if cur.Left == nil && cur.Right == nil {
		return
	}
	*res = append(*res, cur.Val)
	if cur.Left != nil {
		collectLeft(cur.Left, res)
		return
	}
	collectLeft(cur.Right, res)
}

func collectRight(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	if cur.Left == nil && cur.Right == nil {
		return
	}
	if cur.Right != nil {
		collectRight(cur.Right, res)
	} else {
		collectRight(cur.Left, res)
	}
	*res = append(*res, cur.Val)
}

func collectLeaves(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	if cur.Left == nil && cur.Right == nil {
		*res = append(*res, cur.Val)
		return
	}
	collectLeaves(cur.Left, res)
	collectLeaves(cur.Right, res)
}

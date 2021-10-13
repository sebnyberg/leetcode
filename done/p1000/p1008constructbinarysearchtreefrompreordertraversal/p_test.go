package p1008constructbinarysearchtreefrompreordertraversal

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func Test_bstFromPreorder(t *testing.T) {
	for _, tc := range []struct {
		preorder []int
		want     string
	}{
		{[]int{8, 5, 1, 7, 10, 12}, "[8,5,10,1,7,null,12]"},
		{[]int{1, 3}, "[1,null,3]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.preorder), func(t *testing.T) {
			want := ParseTree(tc.want)
			require.Equal(t, want, bstFromPreorder(tc.preorder))
		})
	}
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return &TreeNode{}
	}
	// Approach: put stuff on left side unless doing so would violate the BFS
	// condition. Otherwise, put it on the right side.
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{{Val: math.MaxInt32}, root}
	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{Val: preorder[i]}

		// Pop from stack until the new node can be added.
		parent := stack[len(stack)-1]
		for stack[len(stack)-1].Val < preorder[i] {
			parent = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if node.Val < parent.Val {
			parent.Left = node
		} else {
			parent.Right = node
		}
		stack = append(stack, node)
	}
	return root
}

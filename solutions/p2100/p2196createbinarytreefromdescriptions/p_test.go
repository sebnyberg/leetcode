package p2196createbinarytreefromdescriptions

import (
	"container/list"
	"fmt"
	"leetcode"
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

func Test_createBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		descriptions [][]int
		want         string
	}{
		{leetcode.ParseMatrix("[[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]"), "[50,20,80,15,17,19]"},
		{leetcode.ParseMatrix("[[1,2,1],[2,3,0],[3,4,1]]"), "[1,2,null,null,3,4]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.descriptions), func(t *testing.T) {
			tree := ParseTree(tc.want)
			res := createBinaryTree(tc.descriptions)
			require.True(t, res.Equals(tree))
		})
	}
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	isChild := make(map[int]bool)
	getNode := func(v int) *TreeNode {
		if n, exists := nodes[v]; exists {
			return n
		}
		nodes[v] = &TreeNode{Val: v}
		isChild[v] = false
		return nodes[v]
	}
	for _, descr := range descriptions {
		parent, child, isLeft := descr[0], descr[1], descr[2]
		parentNode := getNode(parent)
		childNode := getNode(child)
		isChild[child] = true
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}
	for k, v := range isChild {
		if !v {
			res := nodes[k]
			return res
		}
	}
	return nil
}

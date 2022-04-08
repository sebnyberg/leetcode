package p0653twosumivinputisabst

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

func Test_constructMaximumBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want [][]string
	}{
		{
			"[1,2]", [][]string{{"", "1", ""}, {"2", "", ""}},
		},
		{
			"[1,2,3,null,4]", [][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tree := ParseTree(tc.tree)
			res := printTree(tree)
			require.Equal(t, tc.want, res)
		})
	}
}

func printTree(root *TreeNode) [][]string {
	height := findHeight(root) - 1
	// For some reason, the exercise considers the height to exclude the root..
	width := (1 << (height + 1)) - 1
	type nodePos struct {
		node *TreeNode
		c    int
	}
	if root == nil {
		return nil
	}
	curr := []nodePos{{root, (width - 1) / 2}}
	next := []nodePos{}
	var res [][]string
	for r := 0; len(curr) > 0; r++ {
		next = next[:0]
		row := make([]string, width)
		for _, n := range curr {
			row[n.c] = fmt.Sprint(n.node.Val)
			if n.node.Left != nil {
				next = append(next, nodePos{n.node.Left, n.c - (1 << (height - r - 1))})
			}
			if n.node.Right != nil {
				next = append(next, nodePos{n.node.Right, n.c + (1 << (height - r - 1))})
			}
		}
		res = append(res, row)

		curr, next = next, curr
	}

	return res
}

func findHeight(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return 1 + max(findHeight(cur.Left), findHeight(cur.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

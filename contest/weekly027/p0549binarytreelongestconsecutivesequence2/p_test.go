package p0549binarytreelongestconsecutivesequence2

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

func Test_longestConsecutive(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want int
	}{
		{"[1,2,null,3,null,4]", 4},
		{"[1,2,3]", 2},
		{"[2,1,3]", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tree := ParseTree(tc.tree)
			require.Equal(t, tc.want, longestConsecutive(tree))
		})
	}
}

func longestConsecutive(root *TreeNode) int {
	var res int
	findLongestConsecutive(root, math.MaxInt32, &res)
	return res
}

func findLongestConsecutive(cur *TreeNode, parent int, res *int) (incr, decr int) {
	if cur == nil {
		return 0, 0
	}

	incrLeft, decrLeft := findLongestConsecutive(cur.Left, cur.Val, res)
	incrRight, decrRight := findLongestConsecutive(cur.Right, cur.Val, res)

	if cur.Val == parent+1 {
		// This is part of an increasing chain
		incr = 1 + max(incrLeft, incrRight)
	} else if cur.Val == parent-1 {
		// This is part of a decreasing chain
		decr = 1 + max(decrLeft, decrRight)
	}

	maxLen := max(
		1+decrLeft+incrRight,
		1+incrLeft+decrRight,
	)
	*res = max(*res, maxLen)

	return incr, decr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

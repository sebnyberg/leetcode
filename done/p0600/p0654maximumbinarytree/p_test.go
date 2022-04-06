package p0654maximumbinarytree

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
		nums []int
		want string
	}{
		{[]int{3, 2, 1, 6, 0, 5}, "[6,3,5,null,2,0,null,null,1]"},
		{[]int{3, 2, 1}, "[3,null,2,null,1]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			tree := ParseTree(tc.want)
			res := constructMaximumBinaryTree(tc.nums)
			require.True(t, res.Equals(tree))
		})
	}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// There aren't that many elements so this can be done brute-force
	if len(nums) == 0 {
		return nil
	}
	i := findMax(nums)
	return &TreeNode{
		Val:   nums[i],
		Left:  constructMaximumBinaryTree(nums[:i]),
		Right: constructMaximumBinaryTree(nums[i+1:]),
	}
}

func findMax(nums []int) int {
	i := -1
	max := math.MinInt32
	for j, n := range nums {
		if n > max {
			max = n
			i = j
		}
	}
	return i
}

package p1339maximumproductofsplitbinarytree

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

func Test_maxProduct(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want int
	}{
		{"[1,2,3,4,5,6]", 110},
		{"[1,null,2,3,4,null,null,5,6]", 90},
		{"[2,3,9,10,7,8,6,5,4,11,1]", 1025},
		{"[1,1]", 1},
		{"[6,10,null,null,6,1,null,1]", 128},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			root := ParseTree(tc.tree)
			require.Equal(t, tc.want, maxProduct(root))
		})
	}
}

func maxProduct(root *TreeNode) int {
	var f maxProductFinder
	f.totalSize = f.visit(root)
	f.visit(root)
	return f.maxProduct % (1e9 + 7)
}

type maxProductFinder struct {
	maxProduct int
	totalSize  int
}

// visit visits the provided node and its children, returning the sum of its
// subtree.
func (f *maxProductFinder) visit(cur *TreeNode) int {
	if cur == nil {
		return 0
	}

	leftSum := f.visit(cur.Left)
	rightSum := f.visit(cur.Right)

	if cur.Left != nil {
		f.maxProduct = max(f.maxProduct, leftSum*(f.totalSize-leftSum))
	}
	if cur.Right != nil {
		f.maxProduct = max(f.maxProduct, rightSum*(f.totalSize-rightSum))
	}
	return leftSum + rightSum + cur.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

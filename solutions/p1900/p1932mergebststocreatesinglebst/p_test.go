package p1932mergebststocreatesinglebst

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

func Test_canMerge(t *testing.T) {
	for _, tc := range []struct {
		trees []string
		want  string
	}{
		{[]string{"[1,null,3]", "[3,1]", "[4,2]"}, "[]"},
		{[]string{"[2,1,3]", "[3,2]"}, "[]"},
		{[]string{"[2,1]", "[3,2,5]", "[5,4]"}, "[3,2,5,1,null,4]"},
		{[]string{"[5,3,8]", "[3,2,6]"}, "[]"},
		{[]string{"[5,4]", "[3]"}, "[]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.trees), func(t *testing.T) {
			trees := make([]*TreeNode, len(tc.trees))
			for i, t := range tc.trees {
				trees[i] = ParseTree(t)
			}
			res := canMerge(trees)
			require.True(t, res.Equals(ParseTree(tc.want)))
		})
	}
}

func canMerge(trees []*TreeNode) *TreeNode {
	indeg := make(map[int]int)
	nodes := make(map[int]*TreeNode)

	// Helper for maybe counting indegrees / adding to global nodes map
	maybeAdd := func(n *TreeNode) {
		if n == nil {
			return
		}
		indeg[n.Val]++
		if _, exists := nodes[n.Val]; !exists {
			nodes[n.Val] = n
		}
	}
	// Count in-degrees and collect nodes into the nodes map
	for _, t := range trees {
		if _, exists := indeg[t.Val]; !exists {
			indeg[t.Val] = 0
		}
		nodes[t.Val] = t // overrides leaf node (if any)
		maybeAdd(t.Left)
		maybeAdd(t.Right)
	}

	// Only one node should have an indegree of zero - the root
	var root *TreeNode
	for v, deg := range indeg {
		if deg == 1 {
			continue
		}
		if root != nil {
			return nil
		}
		root = nodes[v]
	}

	// Perform inorder traversal of nodes, ensuring that each node is the node
	// in the nodes map. This will override leaf node addresses with the roots
	// they should be attached to
	seen := make(map[int]bool)
	min := math.MinInt32
	var inorder func(n **TreeNode) bool
	inorder = func(n **TreeNode) bool {
		if (*n) == nil {
			return true
		}
		seen[(*n).Val] = true

		// Replace current node (which may be a leaf) with 1. root, 2. leaf
		// If it is the root itself, then this is a no-op
		if m, exists := nodes[(*n).Val]; exists {
			*n = m // Replace pointers to this node with the map entry
		}

		if !inorder(&((*n).Left)) {
			return false
		}

		// Check current node against min value
		if (*n).Val <= min {
			return false
		}

		// Update lower bound based on the currently visited node
		// this will work because we are performing in-order traversal
		min = (*n).Val

		return inorder(&(*n).Right)
	}

	if !inorder(&root) || len(seen) != len(nodes) {
		return nil
	}
	return root
}

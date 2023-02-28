package p0652findduplicatesubtrees

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

func Test_findDuplicatesubtrees(t *testing.T) {
	for _, tc := range []struct {
		root string
		want []string
	}{
		{"[1,2,3,4,null,2,4,null,null,4]", []string{"[2,4]", "[4]"}},
		{"[2,1,1]", []string{"[1]"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			res := findDuplicateSubtrees(root)
			require.Equal(t, len(tc.want), len(res))
			for i, r := range res {
				require.True(t, r.Equals(ParseTree(tc.want[i])))
			}
		})
	}
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var treeSize [][]*TreeNode
	count(root, &treeSize)

	// Check pairs of subtrees with the same size
	var res []*TreeNode
	var dup []bool

	for _, nodes := range treeSize {

		// Reset dup. This idiom should be picked up by the compiler to avoid
		// double alloc.
		dup = append(dup[:0], make([]bool, len(nodes))...)

		for i := range nodes {
			if dup[i] {
				continue
			}
			for j := i + 1; j < len(nodes); j++ {
				if equal(nodes[i], nodes[j]) {
					if !dup[i] {
						res = append(res, nodes[i])
					}
					dup[i] = true
					dup[j] = true
				}
			}
		}
	}
	return res
}

func count(curr *TreeNode, treeSizes *[][]*TreeNode) int {
	if curr == nil {
		return 0
	}
	res := count(curr.Left, treeSizes) + count(curr.Right, treeSizes) + 1
	missing := max(0, res-len(*treeSizes)+1)
	*treeSizes = append(*treeSizes, make([][]*TreeNode, missing)...)
	(*treeSizes)[res] = append((*treeSizes)[res], curr)
	return res
}

func equal(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && equal(a.Left, b.Left) && equal(a.Right, b.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

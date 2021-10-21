package p0653twosum4inputisabst

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

func Test_findTarget(t *testing.T) {
	for _, tc := range []struct {
		root string
		k    int
		want bool
	}{
		{"[5,3,6,2,4,null,7]", 9, true},
		{"[5,3,6,2,4,null,7]", 28, false},
		{"[2,1,3]", 4, true},
		{"[2,1,3]", 1, false},
		{"[2,1,3]", 3, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, findTarget(tree, tc.k))
		})
	}
}

func findTarget(root *TreeNode, k int) bool {
	// Feels like cheating, but the O(n) solution is to simply visit all nodes in
	// the tree and check whether a node has been seen which adds up to k
	nodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	seen := make(map[int]struct{})
	for len(nodes) > 0 {
		nextNodes = nextNodes[:0]
		for _, n := range nodes {
			if n == nil {
				continue
			}
			if _, exists := seen[k-n.Val]; exists {
				return true
			}
			seen[n.Val] = struct{}{}
			nextNodes = append(nextNodes, n.Left, n.Right)
		}
		nextNodes, nodes = nodes, nextNodes
	}
	return false
}

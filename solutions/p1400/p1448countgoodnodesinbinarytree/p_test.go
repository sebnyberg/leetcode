package p1448countgoodnodesinbinarytree

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

func Test_goodNodes(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[3,1,4,3,null,1,5]", 4},
		{"[3,3,null,4,2]", 3},
		{"[1]", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, goodNodes(tree))
		})
	}
}

func goodNodes(root *TreeNode) int {
	// There is only one path from any node to the root, so the invariant is
	// simply to pass the max value along the current path to each child node
	var c goodNodeCounter
	c.visit(root, -10001)
	return c.goodNodes
}

type goodNodeCounter struct {
	goodNodes int
}

// Visit visits a node, marking it as good if its value is >= maxVal
func (c *goodNodeCounter) visit(cur *TreeNode, maxVal int) {
	if cur == nil {
		return
	}
	if cur.Val >= maxVal {
		c.goodNodes++
	}
	maxVal = max(maxVal, cur.Val)
	c.visit(cur.Left, maxVal)
	c.visit(cur.Right, maxVal)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0563binarytreetilt

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

func Test_findTilt(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[1,2,3]", 1},
		{"[4,2,9,3,5,null,7]", 15},
		{"[21,7,14,1,1,2,2,3,3]", 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, findTilt(tree))
		})
	}
}

func findTilt(root *TreeNode) int {
	var sum int
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var visit func(*TreeNode) int
	visit = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := visit(node.Left), visit(node.Right)
		sum += abs(left - right)
		return left + right + node.Val
	}
	visit(root)
	return sum
}

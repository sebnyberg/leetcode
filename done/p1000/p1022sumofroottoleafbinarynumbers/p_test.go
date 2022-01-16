package p10222sumofroottoleafbinarynumbers

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

func Test_sumRootToLeaf(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[1,0,1,0,1,0,1]", 22},
		{"[0]", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.Equal(t, tc.want, sumRootToLeaf(root))
		})
	}
}

func sumRootToLeaf(root *TreeNode) int {
	return sum(root, 0)
}

func sum(cur *TreeNode, val int) int {
	if cur == nil {
		return 0
	}
	val += cur.Val
	if cur.Left == nil && cur.Right == nil {
		// this is a leaf
		return val
	}
	return sum(cur.Left, val<<1) + sum(cur.Right, val<<1)
}

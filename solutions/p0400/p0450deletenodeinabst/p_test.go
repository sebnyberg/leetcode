package p0450deletenodeinabst

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

func Test_deleteNode(t *testing.T) {
	for _, tc := range []struct {
		rootStr string
		key     int
		want    string
	}{
		{"[5,3,6,2,4,null,7]", 3, "[5,4,6,2,null,null,7]"},
		{"[5,3,6,2,4,null,7]", 0, "[5,3,6,2,4,null,7]"},
		{"[]", 0, "[]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rootStr), func(t *testing.T) {
			tree := ParseTree(tc.rootStr)
			res := deleteNode(tree, tc.key)
			want := ParseTree(tc.want)
			require.Equal(t, want, res)
		})
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	default: // root.Val == key
		switch {
		case root.Right == nil:
			return root.Left
		case root.Left == nil:
			return root.Right
		}
		// Find left-most child of right side and put it in the current position
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		// Put left-most child in current node position, then delete it
		root.Val = cur.Val
		root.Right = deleteNode(root.Right, cur.Val)
	}
	return root
}

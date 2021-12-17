package p0938rangesumofbst

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

func Test_RangeSumBST(t *testing.T) {
	for _, tc := range []struct {
		root      string
		low, high int
		want      int
	}{
		{"[10,5,15,3,7,null,18]", 7, 15, 32},
		{"[10,5,15,3,7,13,18,1,null,6]", 6, 10, 23},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.Equal(t, tc.want, rangeSumBST(root, tc.low, tc.high))
		})
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return visit(root, low, high)
}

func visit(cur *TreeNode, low, high int) int {
	if cur == nil {
		return 0
	}

	var res int
	if cur.Val >= low && cur.Val <= high {
		res += cur.Val
	}
	if cur.Val > low {
		res += visit(cur.Left, low, high)
	}
	if cur.Val < high {
		res += visit(cur.Right, low, high)
	}
	return res
}

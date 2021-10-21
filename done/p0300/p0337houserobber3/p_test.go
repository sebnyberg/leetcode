package p0337houserobber3

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

func Test_rob(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[3,2,3,null,3,null,1]", 7},
		{"[3,4,5,1,3,null,1]", 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, rob(tree))
		})
	}
}

func rob(root *TreeNode) int {
	robbed, notRobbed := maybeRob(root)
	return max(robbed, notRobbed)
}

// maybeRob returns the maximum cash possible to rob if this node is robbed,
// and if it is not robbed.
func maybeRob(cur *TreeNode) (int, int) {
	if cur == nil {
		return 0, 0
	}
	leftRob, leftNoRob := maybeRob(cur.Left)
	rightRob, rightNoRob := maybeRob(cur.Right)

	// Robbing this node => not possible to rob sub-nodes
	maxWithRob := cur.Val + leftNoRob + rightNoRob

	// Not robbing this node => possible (but not required) to rob sub-nodes
	maxWithoutRob := max(leftNoRob, leftRob) + max(rightNoRob, rightRob)

	return maxWithRob, maxWithoutRob
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0437pathsum3

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

func Test_pathSum(t *testing.T) {
	for _, tc := range []struct {
		root      string
		targetSum int
		want      int
	}{
		{"[1,-2,-3,1,3,-2,null,-1]", 0, 2},
		{"[1,2,6,1,null,-2]", 4, 2},
		{"[10,5,-3,3,2,null,11,3,-2,null,1]", 8, 3},
		{"[5,4,8,11,null,13,4,7,2,null,null,5,1]", 22, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			res := pathSum(root, tc.targetSum)
			require.Equal(t, tc.want, res)
		})
	}
}

func pathSum(root *TreeNode, targetSum int) int {
	parentSums := make(map[int]uint32, 100)
	res := visit(root, parentSums, 0, targetSum)
	return int(res)
}

func visit(node *TreeNode, parentSums map[int]uint32, curr, target int) uint32 {
	if node == nil {
		return 0
	}
	curr += node.Val
	res := parentSums[curr-target]
	if curr == target {
		res++
	}
	parentSums[curr]++
	left := visit(node.Left, parentSums, curr, target)
	right := visit(node.Right, parentSums, curr, target)
	parentSums[curr]--
	return res + left + right
}

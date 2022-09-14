package p1457pseudopalinpath

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

func Test_pseudoPalindromicPaths(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want int
	}{
		{"[2,3,1,3,1,null,1]", 2},
		{"[2,1,1,1,3,null,null,null,null,null,1]", 1},
		{"[9]", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tree := ParseTree(tc.tree)
			require.Equal(t, tc.want, pseudoPalindromicPaths(tree))
		})
	}
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var count [10]int
	return explore(root, count)
}

func explore(node *TreeNode, count [10]int) int {
	if node == nil {
		return 0
	}
	count[node.Val]++
	res := explore(node.Left, count) + explore(node.Right, count)
	calc := func() int {
		var hasOdd bool
		for _, x := range count {
			if x&1 == 0 {
				continue
			}
			if hasOdd {
				return 0
			}
			hasOdd = true
		}
		return 1
	}
	if node.Left == nil && node.Right == nil {
		res += calc()
	}
	count[node.Val]--
	return res
}

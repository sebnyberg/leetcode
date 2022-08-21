package p2385amountoftimeforbinarytreetobeinfected

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func Test_amountOfTime(t *testing.T) {
	for _, tc := range []struct {
		root  string
		start int
		want  int
	}{
		{"[1,5,3,null,4,10,6,9,2]", 3, 4},
		{"[1]", 1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.Equal(t, tc.want, amountOfTime(root, tc.start))
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	// Collect adjacency list for nodes
	adj := make(map[int][]*TreeNode)

	var startNode *TreeNode
	var fillAdj func(curr *TreeNode)
	fillAdj = func(curr *TreeNode) {
		if curr.Val == start {
			startNode = curr
		}
		if curr.Left != nil {
			adj[curr.Val] = append(adj[curr.Val], curr.Left)
			adj[curr.Left.Val] = append(adj[curr.Left.Val], curr)
			fillAdj(curr.Left)
		}
		if curr.Right != nil {
			adj[curr.Val] = append(adj[curr.Val], curr.Right)
			adj[curr.Right.Val] = append(adj[curr.Right.Val], curr)
			fillAdj(curr.Right)
		}
	}
	fillAdj(root)
	if len(adj) == 0 {
		return 0
	}

	infected := make(map[*TreeNode]struct{})
	infected[startNode] = struct{}{}
	curr := []*TreeNode{startNode}
	next := []*TreeNode{}
	var minute int
	for ; len(infected) != len(adj); minute++ {
		next = next[:0]
		for _, node := range curr {
			for _, neiNode := range adj[node.Val] {
				if _, exists := infected[neiNode]; exists {
					continue
				}
				infected[neiNode] = struct{}{}
				next = append(next, neiNode)
			}
		}
		curr, next = next, curr
	}
	return minute
}

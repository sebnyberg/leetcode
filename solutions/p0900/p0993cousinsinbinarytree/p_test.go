package p0993cousinsinbinarytree

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

func Test_isCousins(t *testing.T) {
	for _, tc := range []struct {
		tree string
		x, y int
		want bool
	}{
		{"[1,2,3,4]", 4, 3, false},
		{"[1,2,3,null,4,null,5]", 5, 4, true},
		{"[1,2,3,null,4]", 2, 3, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tree := ParseTree(tc.tree)
			require.Equal(t, tc.want, isCousins(tree, tc.x, tc.y))
		})
	}
}

func isCousins(root *TreeNode, x int, y int) bool {
	var relation [101][2]int8
	gather(root, 0, -1, &relation)
	return relation[x][0] == relation[y][0] && relation[x][1] != relation[y][1]
}

func gather(cur *TreeNode, level, parent int8, relation *[101][2]int8) {
	if cur == nil {
		return
	}
	relation[cur.Val][0] = level
	relation[cur.Val][1] = parent
	gather(cur.Left, level+1, int8(cur.Val), relation)
	gather(cur.Right, level+1, int8(cur.Val), relation)
}

package p0501findmodeinbinarysearchtree

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

func (n *TreeNode) FindVal(val int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val == val {
		return n
	}
	if l := n.Left.FindVal(val); l != nil {
		return l
	}
	return n.Right.FindVal(val)
}

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == nil || other == nil {
		return n == other
	}
	if n.Val != other.Val {
		return false
	}
	return n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
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

func Test_findMode(t *testing.T) {
	for _, tc := range []struct {
		root string
		want []int
	}{
		{"[1,null,2]", []int{1, 2}},
		{"[1,null,2,2]", []int{2}},
		{"[0]", []int{0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.ElementsMatch(t, tc.want, findMode(root))
		})
	}
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int, 1000)
	traverse(root, m)
	var maxCount int
	var res []int
	for v, count := range m {
		if count > maxCount {
			res = res[:0]
			res = append(res, v)
			maxCount = count
		} else if count == maxCount {
			res = append(res, v)
		}
	}
	return res
}

func traverse(cur *TreeNode, m map[int]int) {
	if cur == nil {
		return
	}
	m[cur.Val]++
	traverse(cur.Left, m)
	traverse(cur.Right, m)
}

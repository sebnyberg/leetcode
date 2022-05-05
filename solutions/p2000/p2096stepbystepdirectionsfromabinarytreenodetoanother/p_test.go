package p2096stepbystepdirectionsfromabinarytreenodetoanother

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

func Test_getDirections(t *testing.T) {
	for _, tc := range []struct {
		tree                  string
		startValue, destValue int
		want                  string
	}{
		{"[1,null,10,12,13,4,6,null,15,null,null,5,11,null,2,14,7,null,8,null,null,null,9,3]", 6, 15, "UURR"},
		{"[5,1,2,3,null,6,4]", 3, 6, "UURL"},
		{"[2,1]", 2, 1, "L"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			tree := ParseTree(tc.tree)
			require.Equal(t, tc.want, getDirections(tree, tc.startValue, tc.destValue))
		})
	}
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	prefix := make([]byte, 1e5)
	f := pathFinder{}
	_, _ = f.visit(root, prefix, 0, startValue, destValue)
	res := strings.Repeat("U", f.startLevel-f.rootLevel)
	res += string(f.destPath[f.rootLevel:])
	return res
}

type pathFinder struct {
	startLevel int
	rootLevel  int
	destPath   []byte
}

func (f *pathFinder) visit(cur *TreeNode, path []byte, level, startValue, destValue int) (bool, bool) {
	if cur == nil {
		return false, false
	}
	var hasStart bool
	if cur.Val == startValue {
		f.startLevel = level
		hasStart = true
	}
	var hasEnd bool
	if cur.Val == destValue {
		hasEnd = true
		f.destPath = make([]byte, level)
		copy(f.destPath, path)
	}
	// Go left
	path[level] = 'L'
	leftHasStart, leftHasEnd := f.visit(cur.Left, path, level+1, startValue, destValue)
	path[level] = 'R'
	rightHasStart, rightHasEnd := f.visit(cur.Right, path, level+1, startValue, destValue)
	if leftHasStart && leftHasEnd || rightHasStart && rightHasEnd {
		return true, true
	}
	if leftHasStart && hasEnd || rightHasStart && hasEnd ||
		leftHasStart && rightHasEnd || leftHasEnd && rightHasStart ||
		leftHasEnd && hasStart || rightHasEnd && hasStart {
		f.rootLevel = level
		return true, true
	}
	return leftHasStart || hasStart || rightHasStart, leftHasEnd || hasEnd || rightHasEnd
}

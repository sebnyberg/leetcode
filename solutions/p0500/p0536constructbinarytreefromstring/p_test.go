package p0536constructbinarytreefromstring

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
func Test_str2tree(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"4(2(3)(1))(6(5))", "[4,2,6,3,1,5]"},
		{"4(2(3)(1))(6(5)(7))", "[4,2,6,3,1,5,7]"},
		{"-4(2(3)(1))(6(5)(7))", "[-4,2,6,3,1,5,7]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			want := ParseTree(tc.want)
			res := str2tree(tc.s)
			require.True(t, res.Equals(want))
		})
	}
}

func str2tree(s string) *TreeNode {
	if s == "" {
		return nil
	}
	s = "(" + s + ")"
	stack := []*TreeNode{{}}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(':
			// Push new node to stack
			stack = append(stack, &TreeNode{})
		case ')':
			// Connect node to parent in stack
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if par := stack[len(stack)-1]; par.Left == nil {
				par.Left = cur
			} else {
				par.Right = cur
			}
		default:
			// Parse number
			j := i + 1
			for ; j < len(s) && s[j] != '(' && s[j] != ')'; j++ {
			}
			x, err := strconv.Atoi(s[i:j])
			if err != nil {
				log.Fatalln(err)
			}
			stack[len(stack)-1].Val = x
			i = j - 1
		}
	}
	return stack[0].Left
}

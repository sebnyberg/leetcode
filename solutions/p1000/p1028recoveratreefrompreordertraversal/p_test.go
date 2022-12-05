package p1028recoveratreefrompreordertraversal

import (
	"fmt"
	"testing"
	"unicode"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_recoverFromPreorder(t *testing.T) {
	for i, tc := range []struct {
		traversal string
	}{
		{"1-2--3--4-5--6--7"},
		{"1-2--3---4-5--6---7"},
		{"1-401--349---90--88"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			recoverFromPreorder(tc.traversal)
		})
	}
}

func recoverFromPreorder(traversal string) *TreeNode {
	_, rootVal, i := parse(traversal)
	root := &TreeNode{
		Val: rootVal,
	}
	stack := []*TreeNode{root}
	nchild := []int{0}
	m := 1
	for i < len(traversal) {
		d, x, j := parse(traversal[i:])
		if d <= 0 {
			fmt.Println(traversal[i:])
		}
		// While d < len(stack), pop
		for d < m {
			stack = stack[:m-1]
			nchild = nchild[:m-1]
			m--
		}
		node := &TreeNode{
			Val: x,
		}
		if nchild[m-1] == 0 {
			stack[m-1].Left = node
			nchild[m-1]++
		} else {
			// Add to right
			stack[m-1].Right = node
			nchild[m-1]++
		}
		stack = append(stack, node)
		nchild = append(nchild, 0)
		m++

		i += j
	}
	return root
}

func parse(s string) (d int, x int, i int) {
	for i < len(s) && s[i] == '-' {
		i++
		d++
	}
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		x *= 10
		x += int(s[i] - '0')
		i++
	}
	return d, x, i
}

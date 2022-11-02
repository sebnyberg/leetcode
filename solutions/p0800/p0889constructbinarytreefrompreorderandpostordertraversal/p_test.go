package p0889constructbinarytreefrompreorderandpostordertraversal

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_constructFromPrePost(t *testing.T) {
	for i, tc := range []struct {
		preorder  []int
		postorder []int
		want      *TreeNode
	}{
		{[]int{2, 1}, []int{1, 2}, nil},
		{[]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, nil},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := constructFromPrePost(tc.preorder, tc.postorder)
			_ = res
			// require.Equal(t, tc.want, constructFromPrePost(tc.preorder, tc.postorder))
		})
	}
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	node := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return node
	}
	child := preorder[1]
	// if the child is the only child of the node, then it is at the end of the
	// postorder slice
	if postorder[len(postorder)-2] == child {
		node.Right = constructFromPrePost(preorder[1:], postorder[:len(postorder)-1])
		return node
	}
	left := child
	var j int
	for ; postorder[j] != left; j++ {
	}
	nleft := j + 1
	node.Left = constructFromPrePost(preorder[1:nleft+1], postorder[:nleft])
	node.Right = constructFromPrePost(preorder[1+nleft:], postorder[nleft:len(postorder)-1])
	return node
}

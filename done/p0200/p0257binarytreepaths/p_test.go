package p0257binarytreepaths

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var f pathFinder
	f.findPaths(root, []string{})
	return f.paths
}

type pathFinder struct {
	paths []string
}

func (f *pathFinder) findPaths(cur *TreeNode, vals []string) {
	vals = append(vals, strconv.Itoa(cur.Val))
	if cur.Left == nil && cur.Right == nil {
		f.paths = append(f.paths, strings.Join(vals, "->"))
		return
	}
	if cur.Right != nil {
		f.findPaths(cur.Right, vals)
	}
	if cur.Left != nil {
		if cur.Right != nil {
			valsCpy := make([]string, len(vals))
			copy(valsCpy, vals)
			vals = valsCpy
		}
		f.findPaths(cur.Left, vals)
	}
}

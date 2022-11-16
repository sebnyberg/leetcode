package p0987verticalordertrav

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type entry struct {
	i, j int
	val  int
}

func verticalTraversal(root *TreeNode) [][]int {
	var entries []entry
	visit(root, &entries, 0, 0)
	sort.Slice(entries, func(i, j int) bool {
		a := entries[i]
		b := entries[j]
		if a.j == b.j {
			if a.i == b.i {
				return a.val < b.val
			}
			return a.i < b.i
		}
		return a.j < b.j
	})
	var res [][]int
	for i := 0; i < len(entries); i++ {
		k := len(res)
		res = append(res, []int{})
		res[k] = append(res[k], entries[i].val)
		for i < len(entries)-1 && entries[i+1].j == entries[i].j {
			i++
			res[k] = append(res[k], entries[i].val)
		}
	}
	return res
}

func visit(cur *TreeNode, entries *[]entry, i, j int) {
	if cur == nil {
		return
	}
	*entries = append(*entries, entry{i, j, cur.Val})
	visit(cur.Left, entries, i+1, j-1)
	visit(cur.Right, entries, i+1, j+1)
}

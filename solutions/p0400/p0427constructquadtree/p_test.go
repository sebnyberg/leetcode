package p0427constructquadtree

import (
	"testing"

	"github.com/sebnyberg/leetcode"
)

func TestConstruct(t *testing.T) {
	input := leetcode.ParseMatrix("[[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]")
	res := construct(input)
	_ = res
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	m, n := len(grid), len(grid[0])
	res := parse(grid, 0, 0, m, n)
	return res
}

func parse(grid [][]int, i1, j1, i2, j2 int) *Node {
	cur := new(Node)
	val := grid[i1][j1]
	for i := i1; i < i2; i++ {
		for j := j1; j < j2; j++ {
			if grid[i][j] != val {
				w := (i2 - i1) / 2
				cur.TopLeft = parse(grid, i1, j1, i1+w, j1+w)
				cur.TopRight = parse(grid, i1, j1+w, i1+w, j1+w+w)
				cur.BottomLeft = parse(grid, i1+w, j1, i1+w+w, j1+w)
				cur.BottomRight = parse(grid, i1+w, j1+w, i1+w+w, j1+w+w)
				return cur
			}
		}
	}
	cur.IsLeaf = true
	cur.Val = val == 1
	return cur
}

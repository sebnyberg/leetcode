package p0314binarytreeverticalordertraversal

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_verticalOrder(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want [][]int
	}{
		{
			"[0,1,4,2,8,37,13,3,34,14,29,66,45,22,19,5,6,39,69,null,17,null,35,null,null,null,null,32,null,null,58,null,9,10,7,null,55,89,null,42,51,57,null,86,null,null,null,11,18,53,15,12,null,null,null,null,null,48,null,80,84,75,65,null,null,26,64,27,21,9]",
			[][]int{{26, 9}, {5, 11, 53}, {3, 9, 10, 64, 27}, {2, 6, 39, 18, 15, 12, 48}, {1, 34, 14, 66, 7, 55, 89, 42, 86, 21}, {0, 8, 37, 69, 17, 32, 80, 75}, {4, 29, 45, 22, 51, 57}, {13, 35, 84, 65}, {19}, {58}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			root := ParseTree(tc.tree)
			require.Equal(t, tc.want, verticalOrder(root))
		})
	}
}

func verticalOrder(root *TreeNode) [][]int {
	coll := nodeCollector{
		values:      make([][][3]int, 250),
		minVertical: math.MaxInt32,
		maxVertical: math.MinInt32,
	}
	if root == nil {
		return nil
	}
	coll.visit(root, 0, 100)
	coll.values = coll.values[coll.minVertical : coll.maxVertical+1]
	res := make([][]int, len(coll.values))
	for i, vals := range coll.values {
		sort.Slice(vals, func(i, j int) bool {
			if vals[i][1] == vals[j][1] {
				return vals[i][2] < vals[j][2]
			}
			return vals[i][1] < vals[j][1]
		})
		res[i] = make([]int, len(vals))
		for j := range vals {
			res[i][j] = vals[j][0]
		}
	}
	return res
}

type nodeCollector struct {
	values                   [][][3]int
	minVertical, maxVertical int
}

func (c *nodeCollector) visit(cur *TreeNode, level, vertical int) {
	if cur == nil {
		return
	}
	c.minVertical = min(c.minVertical, vertical)
	c.maxVertical = max(c.maxVertical, vertical)
	c.visit(cur.Left, level+1, vertical-1)
	c.values[vertical] = append(c.values[vertical], [3]int{cur.Val, level, len(c.values[vertical])})
	c.visit(cur.Right, level+1, vertical+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

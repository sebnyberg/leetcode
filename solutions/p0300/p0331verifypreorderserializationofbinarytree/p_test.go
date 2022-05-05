package p0331verifypreorderserializationofbinarytree

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValidSerialization(t *testing.T) {
	for _, tc := range []struct {
		preorder string
		want     bool
	}{
		{"9,3,4,#,#,1,#,#,2,#,6,#,#", true},
		{"1,#", false},
		{"9,#,#,1", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.preorder), func(t *testing.T) {
			require.Equal(t, tc.want, isValidSerialization(tc.preorder))
		})
	}
}

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	nvisited := visit(nodes, 0, len(nodes))
	return nvisited+1 == len(nodes)
}

// Visit node[i], returning the maximum reachable index from that node.
func visit(nodes []string, i, n int) int {
	if i >= n {
		return n
	}
	if nodes[i] == "#" {
		return i
	}
	left := visit(nodes, i+1, n)
	return visit(nodes, left+1, n)
}

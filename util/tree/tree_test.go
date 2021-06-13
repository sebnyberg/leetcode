package tree

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_NewTreeFromList(t *testing.T) {
// 	for _, tc := range []struct {
// 		in   []interface{}
// 		want *TreeNode
// 	}{
// 		{[]interface{}{1, nil}, &TreeNode{Val: 1}},
// 		{[]interface{}{1, 2, nil}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
// 		{[]interface{}{1, 2, nil, 3, nil, nil, 4}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
// 			require.Equal(t, tc.want, NewTreeFromList(tc.in))
// 		})
// 	}
// }

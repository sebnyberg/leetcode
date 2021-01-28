package p1613pathwithmineffort

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_minimumEffortPath(t *testing.T) {
// 	for _, tc := range []struct {
// 		heights [][]int
// 		want    int
// 	}{
// 		{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2},
// 		{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1},
// 		{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}, 0},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.heights), func(t *testing.T) {
// 			require.Equal(t, tc.want, minimumEffortPath(tc.heights))
// 		})
// 	}
// }

// func minimumEffortPath(heights [][]int) int {
// 	// Plan:
// 	// Iterate over all entries in heights
// }

package p0200numberofislands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numIslands(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		want int
	}{
		// {
		// 	[][]byte{
		// 		{'1', '1', '1', '1', '0'},
		// 		{'1', '1', '0', '1', '0'},
		// 		{'1', '1', '0', '0', '0'},
		// 		{'0', '0', '0', '0', '0'},
		// 	},
		// 	1,
		// },
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, numIslands(tc.grid))
		})
	}
}

func numIslands(grid [][]byte) int {
	return 0
}

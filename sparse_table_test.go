package leetcode_test

import (
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSparseTable(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		table := leetcode.NewSparseTable([]int{1, 4, 6, 4, 2, 6, 4, 3},
			func(a, b int) int {
				if a < b {
					return a
				}
				return b
			},
		)
		require.Equal(t, 1, table.Query(0, 3))
		require.Equal(t, 2, table.Query(3, 4))
		require.Equal(t, 2, table.Query(1, 8))
		require.Equal(t, 4, table.Query(3, 3))
		require.Equal(t, 2, table.Query(4, 5))
		require.Equal(t, 1, table.Query(0, 1000))
	})
}

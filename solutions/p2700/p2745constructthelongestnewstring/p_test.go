package p2745constructthelongestnewstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestString(t *testing.T) {
	for i, tc := range []struct {
		x    int
		y    int
		z    int
		want int
	}{
		{2, 5, 1, 12},
		{3, 2, 2, 14},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestString(tc.x, tc.y, tc.z))
		})
	}
}

func longestString(x int, y int, z int) int {
	// AA -> BB
	// BB -> (AA or AB)
	// AB -> (AA or AB)
	//
	// It seems like we can always weave AAs and BBs no matter if we have ABs
	//
	// And we can get rid of all ABs in a row.
	// And it seems ABs are always possible to get rid of and they change
	// nothing really.
	res := (min(x, y)*2 + min(1, abs(x-y)) + z) * 2
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

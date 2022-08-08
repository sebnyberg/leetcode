package p0780reachingpoints

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reachingPoints(t *testing.T) {
	for _, tc := range []struct {
		sx, sy, tx, ty int
		want           bool
	}{
		{1, 1, 3, 5, true},
		{1, 1, 2, 2, false},
		{1, 1, 1, 1, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sx), func(t *testing.T) {
			require.Equal(t, tc.want, reachingPoints(tc.sx, tc.sy, tc.tx, tc.ty))
		})
	}
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if ty < tx {
		tx, ty = ty, tx
		sx, sy = sy, sx
	}
	for {
		if tx < sx {
			return false
		}
		if tx == sx {
			if ty >= sy {
				return ((ty - sy) % sx) == 0
			}
			return false
		}
		tx, ty = ty%tx, tx
		sx, sy = sy, sx
	}
}

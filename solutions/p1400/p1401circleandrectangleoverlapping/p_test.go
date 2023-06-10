package p1401circleandrectangleoverlapping

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkOverlap(t *testing.T) {
	for i, tc := range []struct {
		radius                           int
		xCenter, yCenter, x1, y1, x2, y2 int
		want                             bool
	}{
		{1, 1, 1, 1, -3, 2, -1, false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, checkOverlap(tc.radius, tc.xCenter, tc.yCenter, tc.x1, tc.y1, tc.x2, tc.y2))
		})
	}
}

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	var dx int
	if xCenter < x1 {
		dx = x1 - xCenter
	} else if xCenter > x2 {
		dx = xCenter - x2
	}
	var dy int
	if yCenter < y1 {
		dy = y1 - yCenter
	} else if yCenter > y2 {
		dy = yCenter - y2
	}
	return dx*dx+dy*dy <= radius*radius
}

package p2184numberofwaystobuildsturdybrickwall

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildWall(t *testing.T) {
	for _, tc := range []struct {
		height int
		width  int
		bricks []int
		want   int
	}{
		{2, 3, []int{1, 2}, 2},
		{1, 1, []int{5}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.height), func(t *testing.T) {
			require.Equal(t, tc.want, buildWall(tc.height, tc.width, tc.bricks))
		})
	}
}

func buildWall(height int, width int, bricks []int) int {
	// skip any bricks wider than width
	var j int
	for i := 0; i < len(bricks); i++ {
		if bricks[i] <= width {
			bricks[j] = bricks[i]
			j++
		}
	}
	bricks = bricks[:j]
	if len(bricks) == 0 {
		return 0
	}
	configs := getConfigs(bricks, 0, len(bricks), width)
}

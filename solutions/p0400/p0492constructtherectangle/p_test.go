package p0492constructtherectangle

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constructRectangle(t *testing.T) {
	for _, tc := range []struct {
		area int
		want []int
	}{
		{122122, []int{427, 286}},
		{4, []int{2, 2}},
		{37, []int{37, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.area), func(t *testing.T) {
			require.Equal(t, tc.want, constructRectangle(tc.area))
		})
	}
}

func constructRectangle(area int) []int {
	for width := int(math.Sqrt(float64(area))); width >= 1; width-- {
		if area%width == 0 {
			return []int{area / width, width}
		}
	}
	return nil
}

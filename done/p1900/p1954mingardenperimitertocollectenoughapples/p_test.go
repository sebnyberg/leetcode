package p1954mingardenperimitertocollectenoughapples

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumPerimeter(t *testing.T) {
	for _, tc := range []struct {
		neededApples int64
		want         int64
	}{
		{61, 24},
		{1, 8},
		{13, 16},
		{17, 16},
		{1000000000, 5040},
	} {
		t.Run(fmt.Sprintf("%+v", tc.neededApples), func(t *testing.T) {
			require.Equal(t, tc.want, minimumPerimeter(tc.neededApples))
		})
	}
}

func minimumPerimeter(neededApples int64) int64 {
	sideLen := 3
	sideVal := 5
	radius := 1
	val := 4*sideVal - 4*2*radius
	for val < int(neededApples) {
		radius++
		sideVal = sideVal + sideLen + 2*2*radius
		sideLen += 2
		val += 4*sideVal - 4*2*radius
	}
	return int64(sideLen*4 - 4)
}

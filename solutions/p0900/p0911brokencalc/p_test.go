package p0911brokencalc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_brokenCalc(t *testing.T) {
	for _, tc := range []struct {
		X    int
		Y    int
		want int
	}{
		{1, 1000000000, 39},
		// {2, 3, 2},
		// {5, 8, 2},
		// {3, 10, 3},
		// {1024, 1, 1023},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.X, tc.Y), func(t *testing.T) {
			require.Equal(t, tc.want, brokenCalc(tc.X, tc.Y))
		})
	}
}

func brokenCalc(X int, Y int) (ops int) {
	got := X
	want := Y
	for want > got {
		if want%2 == 1 {
			want++
		} else {
			want /= 2
		}
		ops++
	}
	if got > want {
		ops += got - want
	}
	return ops
}

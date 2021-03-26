package p0223rectanglearea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_computeArea(t *testing.T) {
	for _, tc := range []struct {
		a, b, c, d, e, f, g, h int
		want                   int
	}{
		{-3, 0, 3, 4, 0, -1, 9, 2, 45},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, computeArea(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f, tc.g, tc.h))
		})
	}
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	xOverlap := overlap([2]int{A, C}, [2]int{E, G})
	yOverlap := overlap([2]int{B, D}, [2]int{F, H})
	return (C-A)*(D-B) + (G-E)*(H-F) - xOverlap*yOverlap
}

func overlap(a, b [2]int) int {
	if a[0] > b[0] {
		a, b = b, a
	}
	switch {
	case b[0] > a[1]:
		return 0
	case b[1] < a[1]:
		return b[1] - b[0]
	default:
		return a[1] - b[0]
	}
}

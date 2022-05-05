package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_leastBricks(t *testing.T) {
	for _, tc := range []struct {
		wall [][]int
		want int
	}{
		{[][]int{{100000000}}, 1},
		{[][]int{{1}, {1}, {1}}, 3},
		{[][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.wall), func(t *testing.T) {
			require.Equal(t, tc.want, leastBricks(tc.wall))
		})
	}
}

func Test_findEdges(t *testing.T) {
	for _, tc := range []struct {
		brickRow []int
		want     []int
	}{
		{[]int{1, 2, 2, 1}, []int{0, 2, 4}},
		{[]int{3, 1, 2}, []int{2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.brickRow), func(t *testing.T) {
			e := edgeFinder{buf: make([]int, 0)}
			require.Equal(t, tc.want, e.findEdges(tc.brickRow))
		})
	}
}

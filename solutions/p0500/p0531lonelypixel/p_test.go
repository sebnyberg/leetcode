package p0531lonelypixel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLonelyPixel(t *testing.T) {
	for _, tc := range []struct {
		picture [][]byte
		want    int
	}{
		{[][]byte{
			[]byte("WWB"),
			[]byte("WBW"),
			[]byte("BWW"),
		}, 3},
		{[][]byte{
			[]byte("BBB"),
			[]byte("BBW"),
			[]byte("BBB"),
		}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.picture), func(t *testing.T) {
			require.Equal(t, tc.want, findLonelyPixel(tc.picture))
		})
	}
}

func findLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i := range picture {
		for j, cell := range picture[i] {
			if cell == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}
	var res int
	for i := range picture {
		for j, cell := range picture[i] {
			if cell == 'B' && cols[j] == 1 && rows[i] == 1 {
				res++
			}
		}
	}
	return res
}

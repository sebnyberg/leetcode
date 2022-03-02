package p0533lonelypixel2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findBlackPixel(t *testing.T) {
	for _, tc := range []struct {
		picture [][]byte
		target  int
		want    int
	}{
		{
			[][]byte{
				[]byte("WWBWBWBWWBBW"),
				[]byte("WBWBWWBWWBBW"),
				[]byte("WWWBWWBBBWWW"),
				[]byte("WWBWBWBWWBBW"),
				[]byte("BWWBWWBBBWWW"),
				[]byte("WBWBWWBWWBBW"),
				[]byte("WWBWBWBWWBBW"),
				[]byte("WWBWBWBWWBBW"),
				[]byte("WBWBWWBWWBBW"),
				[]byte("WWWBWWBWWBBW"),
				[]byte("BWWBWWBBBWWW"),
				[]byte("WWWWBWBWWBBW"),
				[]byte("BWWBWWBBBWWW"),
				[]byte("BWWBWWBBBWWW"),
				[]byte("WBWBWWBWWBBW"),
			},
			5, 0,
		},

		{
			[][]byte{
				[]byte("WBWBBW"),
				[]byte("BWBWWB"),
				[]byte("WBWBBW"),
				[]byte("BWBWWB"),
				[]byte("WWWBBW"),
				[]byte("BWBWWB"),
			},
			3,
			9,
		},
		{[][]byte{
			[]byte("WBWBBW"),
			[]byte("WBWBBW"),
			[]byte("WBWBBW"),
			[]byte("WWBWBW"),
		}, 3, 6},
		{[][]byte{
			[]byte("WWB"),
			[]byte("WWB"),
			[]byte("WWB"),
		}, 1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.picture), func(t *testing.T) {
			require.Equal(t, tc.want, findBlackPixel(tc.picture, tc.target))
		})
	}
}

func findBlackPixel(picture [][]byte, target int) int {
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
			if cell != 'B' || cols[j] != target || rows[i] != target {
				continue
			}
			// Check so all rows contain the same row
			ok := true
			for k := 0; k < m; k++ {
				if picture[k][j] != 'B' {
					continue
				}
				if string(picture[k]) != string(picture[i]) {
					ok = false
					break
				}
			}
			if ok {
				res++
			}
		}
	}
	return res
}

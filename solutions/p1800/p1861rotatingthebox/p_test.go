package p1861rotatingthebox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotateTheBox(t *testing.T) {
	for _, tc := range []struct {
		box  [][]byte
		want [][]byte
	}{
		{
			[][]byte{
				{'#', '#', '*', '.', '*', '.'},
				{'#', '#', '#', '*', '.', '.'},
				{'#', '#', '#', '.', '#', '.'},
			},
			[][]byte{
				{'.', '#', '#'},
				{'.', '#', '#'},
				{'#', '#', '*'},
				{'#', '*', '.'},
				{'#', '.', '*'},
				{'#', '.', '.'},
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.box), func(t *testing.T) {
			res := rotateTheBox(tc.box)
			for _, row := range res {
				for _, cell := range row {
					fmt.Print(string(cell))
				}
				fmt.Print("\n")
			}
			require.Equal(t, tc.want, rotateTheBox(tc.box))
		})
	}
}

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
	}
	for i, row := range box {
		var count int
		for j, val := range row {
			switch val {
			case '#':
				res[j][m-1-i] = '.'
				count++
			case '.':
				res[j][m-1-i] = '.'
				continue
			case '*':
				// Write stones to output
				for k := 0; k < count; k++ {
					res[j-1-k][m-1-i] = '#'
				}
				res[j][m-1-i] = '*'
				count = 0
			}
		}
		if count > 0 {
			// Write stones to output
			for k := 0; k < count; k++ {
				res[n-1-k][m-1-i] = '#'
			}
		}
	}
	return res
}

package p1931paintingagridwiththreedifferentcolors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_colorTheGrid(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		want int
	}{
		{1, 2, 6},
		{1, 1, 3},
		{5, 5, 580986},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, colorTheGrid(tc.m, tc.n))
		})
	}
}

func colorTheGrid(m int, n int) int {
	// As usual, let's start with the top-down case.

	// Whether or not a cell can be colored with a certain color depends on:
	// The cell to its left,
	// The cell above.
	//
	// So the number of ways in which a cell can be colored with a certain color
	// is the number of ways in which the cell above was not that color, and the
	// cell to the left was also not that color.
	var mem [1000][1 << 10]int

	res := visitCell(&mem, 0, 0, m, n, 0, 0)

	return res
}

type column int

func (c column) colorAtRow(row int) int {
	// Each color takes up two bits
	// Shift by 2*i to remove all postfix bits, then mask with & 3 to remove pre
	return int(c>>(row*2)) & 3
}

func (c column) setColor(row, color int) column {
	return c | column(color<<(row*2))
}

const bigPrime = 1_000_000_007

func visitCell(mem *[1000][1 << 10]int, row, col, m, n int, curCol, prevCol column) int {
	if row == m { // shift to next column
		return visitCell(mem, 0, col+1, m, n, 0, curCol)
	}
	if col == n {
		return 1
	}
	// At the start of each row, consider whether this state has been visited
	// before
	if row == 0 && mem[col][prevCol] > 0 {
		return mem[col][prevCol]
	}

	var colorAbove int
	if row > 0 {
		colorAbove = curCol.colorAtRow(row - 1)
	}
	colorLeft := prevCol.colorAtRow(row)
	var res int
	for color := 1; color <= 3; color++ {
		if color != colorAbove && color != colorLeft {
			res += visitCell(mem, row+1, col, m, n, curCol.setColor(row, color), prevCol)
			res %= bigPrime
		}
	}
	if row == 0 {
		mem[col][prevCol] = res
	}
	return res
}

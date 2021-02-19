package p0118pascalstriangle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	for _, tc := range []struct {
		numRows int
		want    [][]int
	}{
		{0, [][]int{}},
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numRows), func(t *testing.T) {
			require.Equal(t, tc.want, generate(tc.numRows))
		})
	}
}

func generate(numRows int) [][]int {
	t := PascalsTriangle{
		rows:  [][]int{{1}, {1, 1}},
		nrows: 2,
	}
	for i := 2; i < numRows; i++ {
		t.generateRow()
	}
	return t.rows[:numRows]
}

type PascalsTriangle struct {
	rows  [][]int
	nrows int
}

func (t *PascalsTriangle) generateRow() {
	t.rows = append(t.rows, make([]int, t.nrows+1))
	t.rows[t.nrows][0] = 1
	t.rows[t.nrows][t.nrows] = 1
	for i := 1; i < t.nrows; i++ {
		t.rows[t.nrows][i] = t.rows[t.nrows-1][i-1] + t.rows[t.nrows-1][i]
	}
	t.nrows++
}

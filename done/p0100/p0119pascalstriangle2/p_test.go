package p0119pascalstriangle2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	for _, tc := range []struct {
		numRows int
		want    []int
	}{
		{3, []int{1, 3, 3, 1}},
		{2, []int{1, 2, 1}},
		{1, []int{1, 1}},
		{0, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numRows), func(t *testing.T) {
			require.Equal(t, tc.want, getRow(tc.numRows))
		})
	}
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	row := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		for j := len(row) - 1; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
		row = append(row, 1)
	}
	return row
}

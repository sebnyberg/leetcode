package p2303calculateamountpaidintaxes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateTax(t *testing.T) {
	for _, tc := range []struct {
		brackets [][]int
		income   int
		want     float64
	}{
		{[][]int{{3, 50}, {7, 10}, {12, 25}}, 10, 2.65},
	} {
		t.Run(fmt.Sprintf("%+v", tc.brackets), func(t *testing.T) {
			require.Equal(t, tc.want, calculateTax(tc.brackets, tc.income))
		})
	}
}

func calculateTax(brackets [][]int, income int) float64 {
	var i int
	var res float64
	for income > 0 {
		d := brackets[i][0]
		if i > 0 {
			d -= brackets[i-1][0]
		}
		taxable := min(income, d)
		res += float64(taxable) * (float64(brackets[i][1]) / 100)
		income -= taxable
		i++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

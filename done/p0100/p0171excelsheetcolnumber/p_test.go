package p0171excelsheetcolnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_titleToNumber(t *testing.T) {
	for _, tc := range []struct {
		columnTitle string
		want        int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
		{"FXSHRXW", 2147483647},
	} {
		t.Run(fmt.Sprintf("%+v", tc.columnTitle), func(t *testing.T) {
			require.Equal(t, tc.want, titleToNumber(tc.columnTitle))
		})
	}
}

func titleToNumber(columnTitle string) (res int) {
	n := len(columnTitle)
	multiplier := 1
	for i := range columnTitle {
		a := (columnTitle[n-i-1]-'A')%26 + 1
		res += multiplier * int(a)
		multiplier *= 26
	}
	return res
}

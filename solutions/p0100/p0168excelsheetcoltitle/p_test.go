package p0168excelsheetcoltitle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertToTitle(t *testing.T) {
	for _, tc := range []struct {
		columnNumber int
		want         string
	}{
		// {1, "A"},
		// {28, "AB"},
		// {701, "ZY"},
		{2147483647, "FXSHRXW"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.columnNumber), func(t *testing.T) {
			require.Equal(t, tc.want, convertToTitle(tc.columnNumber))
		})
	}
}

func convertToTitle(columnNumber int) string {
	vals := make([]byte, 0)
	for columnNumber > 0 {
		columnNumber--
		vals = append(vals, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	return string(vals)
}

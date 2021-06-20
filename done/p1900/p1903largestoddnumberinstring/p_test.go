package p1903largestoddnumberinstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestOddNumber(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want string
	}{
		{"52", "5"},
		{"4206", ""},
		{"35427", "35427"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, largestOddNumber(tc.num))
		})
	}
}

func largestOddNumber(num string) string {
	n := len(num)
	for i := n - 1; i >= 0; i-- {
		if (num[i]-'0')%2 == 1 {
			return num[0 : i+1]
		}
	}
	return ""
}

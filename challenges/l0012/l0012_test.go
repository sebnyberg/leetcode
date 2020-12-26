package l0012_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intToRoman(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want string
	}{
		{1000, "M"},
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, intToRoman(tc.in))
		})
	}
}

var romans = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanStrs = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var c int
	var sb strings.Builder
	for i := 0; num > 0 && i < len(romans); i++ {
		c = num / romans[i]
		if c > 0 {
			for j := 0; j < c; j++ {
				sb.WriteString(romanStrs[i])
			}
			num %= romans[i]
		}
	}
	return sb.String()
}

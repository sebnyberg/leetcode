package p2566maximumdifferencebyremappingadigit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMaxDifference(t *testing.T) {
	for i, tc := range []struct {
		num  int
		want int
	}{
		{90, 99},
		{11891, 99009},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minMaxDifference(tc.num))
		})
	}
}

func minMaxDifference(num int) int {
	s := fmt.Sprint(num)
	var a int
	var b int
	var m byte
	var x byte
	for i := range s {
		a *= 10
		if s[i] != '9' && (m == 0 || s[i] == m) {
			m = s[i]
			a += 9
		} else {
			a += int(s[i] - '0')
		}

		b *= 10
		if s[i] != '0' && (x == 0 || s[i] == x) {
			x = s[i]
		} else {
			b += int(s[i] - '0')
		}
	}
	return a - b
}

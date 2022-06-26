package p0738monotoneincreasingdigits

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_monotoneIncreasingDigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{100, 99},
		{10, 9},
		{1234, 1234},
		{332, 299},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, monotoneIncreasingDigits(tc.n))
		})
	}
}

func monotoneIncreasingDigits(n int) int {
	s := []byte(fmt.Sprint(n))
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] < s[i-1] {
			s[i] = '9'
			s[i-1]--
			for j := i + 1; j < len(s); j++ {
				s[j] = '9'
			}
		}
	}
	x, _ := strconv.Atoi(string(s))
	return x
}

package p2843countsymmetricintegers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSymmetricIntegers(t *testing.T) {
	for _, tc := range []struct {
		low  int
		high int
		want int
	}{
		// {1, 100, 9},
		{1200, 1230, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.low), func(t *testing.T) {
			require.Equal(t, tc.want, countSymmetricIntegers(tc.low, tc.high))
		})
	}
}

func countSymmetricIntegers(low int, high int) int {
	var res int
	for x := low; x <= high; x++ {
		s := fmt.Sprint(x)
		n := len(s)
		if n&1 == 1 {
			continue
		}
		var sum int
		for _, b := range s[:n/2] {
			sum += int(b - '0')
		}
		for _, b := range s[n/2:] {
			sum -= int(b - '0')
		}
		if sum == 0 {
			res++
		}
	}
	return res
}

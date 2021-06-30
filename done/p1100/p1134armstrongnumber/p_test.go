package p1134armstrongnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isArmstrong(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{153, true},
		{123, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isArmstrong(tc.n))
		})
	}
}

func isArmstrong(n int) bool {
	var res int
	var count int
	n2 := n
	for n2 > 0 {
		count++
		n2 /= 10
	}
	n2 = n
	for n2 > 0 {
		res += pow(n2%10, count)
		n2 /= 10
	}
	return res == n
}

func pow(a, b int) int {
	c := a
	for b > 1 {
		c *= a
		b--
	}
	return c
}

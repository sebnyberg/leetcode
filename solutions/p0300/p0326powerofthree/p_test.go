package p0326powerofthree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfThree(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{3, true},
		{2, false},
		{0, false},
		{1, true},
		{-1, false},
		{-3, false},
		{-9, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isPowerOfThree(tc.n))
		})
	}
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	ops := 0
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
		ops++
	}
	return true
}

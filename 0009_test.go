package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_l0009(t *testing.T) {
	tcs := []struct {
		in   int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, l0009(tc.in))
		})
	}
}

func l0009(n int) bool {
	if n < 0 {
		return false
	}

	if n < 10 {
		return true
	}

	ns := make([]int8, 0, 32)

	// Put all numbers in a list
	for {
		ns = append(ns, int8(n%10))
		if n < 10 {
			break
		}
		n /= 10
	}

	// Compare numbers
	for i := 0; i < len(ns); i++ {
		if ns[i] != ns[len(ns)-1-i] {
			return false
		}
	}

	return true
}

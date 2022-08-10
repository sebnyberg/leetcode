package p0796rotatestring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotateString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		goal string
		want bool
	}{
		{"abcde", "cdeab", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, rotateString(tc.s, tc.goal))
		})
	}
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	m := len(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[(i+j)%m] != goal[j] {
				break
			}
			if j == len(s)-1 {
				return true
			}
		}
	}
	return false
}

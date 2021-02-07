package p0058lenlastword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLastWord(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{" ", 0},
		{"a ", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLastWord(tc.s))
		})
	}
}

func lengthOfLastWord(s string) (sum int) {
	// Skip spaces
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	// Count non-spaces
	for ; i >= 0 && s[i] != ' '; i-- {
		sum++
	}
	return sum
}

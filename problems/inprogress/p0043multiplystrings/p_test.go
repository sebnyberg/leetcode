package p0043multiplystrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiply(t *testing.T) {
	for _, tc := range []struct {
		num1 string
		num2 string
		want string
	}{
		// {"2", "3", "6"},
		// {"123", "456", "56088"},
	} {
		t.Run(fmt.Sprintf("%v+%v", tc.num1, tc.num2), func(t *testing.T) {
			require.Equal(t, tc.want, multiply(tc.num1, tc.num2))
		})
	}
}

func multiply(num1 string, num2 string) string {
	return ""
}

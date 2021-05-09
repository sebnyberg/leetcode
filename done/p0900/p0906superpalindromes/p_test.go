package p0906superpalindromes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_superpalindromesInRange(t *testing.T) {
	for _, tc := range []struct {
		left  string
		right string
		want  int
	}{
		{"4", "1000", 4},
		{"1", "2", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, superpalindromesInRange(tc.left, tc.right))
		})
	}
}

func superpalindromesInRange(left string, right string) int {
	return 0
}

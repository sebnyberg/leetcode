package p0030substringconcatwords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSubstring(t *testing.T) {
	for _, tc := range []struct {
		s     string
		words []string
		want  []int
	}{
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findSubstring(tc.s, tc.words))
		})
	}
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	return res
}

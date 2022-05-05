package p0454numberofsegmentsinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSegments(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"Hello, my name is John", 5},
		{"Hello", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countSegments(tc.s))
		})
	}
}

func countSegments(s string) int {
	var i int
	var count int
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		count++
		for i < len(s) && s[i] != ' ' {
			i++
		}
	}
	return count
}

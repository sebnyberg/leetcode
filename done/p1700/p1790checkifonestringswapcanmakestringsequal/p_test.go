package p1790checkifonestringswapcanmakestringsequal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areAlmostEqual(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want bool
	}{
		{"bank", "kanb", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, areAlmostEqual(tc.s1, tc.s2))
		})
	}
}

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var nerrs int
	var expect [2]byte
	for i := range s1 {
		if s1[i] != s2[i] {
			if nerrs == 0 {
				expect[0] = s2[i]
				expect[1] = s1[i]
				nerrs++
				continue
			}
			if nerrs == 1 {
				if s1[i] != expect[0] || s2[i] != expect[1] {
					return false
				}
				nerrs++
				continue
			}
			return false
		}
	}
	return nerrs != 1
}

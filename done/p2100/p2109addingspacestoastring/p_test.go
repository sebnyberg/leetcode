package p2109addingspacestoastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addSpaces(t *testing.T) {
	for _, tc := range []struct {
		s      string
		spaces []int
		want   string
	}{
		{"LeetcodeHelpsMeLearn", []int{8, 13, 15}, "Leetcode Helps Me Learn"},
		{"icodeinpython", []int{1, 5, 7, 9}, "i code in py thon"},
		{"spacing", []int{0, 1, 2, 3, 4, 5, 6}, " s p a c i n g"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, addSpaces(tc.s, tc.spaces))
		})
	}
}

func addSpaces(s string, spaces []int) string {
	res := make([]byte, 0, len(s))
	var j int
	for i := range s {
		if j < len(spaces) && i == spaces[j] {
			res = append(res, ' ')
			j++
		}
		res = append(res, s[i])
	}
	return string(res)
}
